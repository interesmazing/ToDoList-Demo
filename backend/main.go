package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// Todo struct represents a single todo item
type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var db *sql.DB

func main() {
	// Database connection
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Ping the database to ensure connectivity
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}
	fmt.Println("Successfully connected to the database!")

	// Set up Gin router
	router := gin.Default()

	// CORS configuration
	config := cors.DefaultConfig()
	// Allow all origins, but in a real production app, you should restrict this
	// to your frontend's domain.
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept"}
	router.Use(cors.New(config))

	// API routes
	api := router.Group("/api")
	{
		api.GET("/todos", getTodos)
		api.POST("/todos", createTodo)
		api.PUT("/todos/:id", updateTodo)
		api.DELETE("/todos/:id", deleteTodo)
	}

	// Start the server
	log.Println("Starting server on port 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

// getTodos handles GET /api/todos
func getTodos(c *gin.Context) {
	rows, err := db.Query("SELECT id, title, completed FROM todos ORDER BY created_at DESC")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve todos"})
		return
	}
	defer rows.Close()

	todos := []Todo{}
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan todo"})
			return
		}
		todos = append(todos, todo)
	}

	c.JSON(http.StatusOK, todos)
}

// createTodo handles POST /api/todos
func createTodo(c *gin.Context) {
	var newTodo struct {
		Title string `json:"title"`
	}

	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if newTodo.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title cannot be empty"})
		return
	}

	var createdTodo Todo
	err := db.QueryRow(
		"INSERT INTO todos (title) VALUES ($1) RETURNING id, title, completed",
		newTodo.Title,
	).Scan(&createdTodo.ID, &createdTodo.Title, &createdTodo.Completed)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	c.JSON(http.StatusCreated, createdTodo)
}

// updateTodo handles PUT /api/todos/:id
func updateTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updates struct {
		Title     *string `json:"title"`
		Completed *bool   `json:"completed"`
	}

	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Check if there is anything to update
	if updates.Title == nil && updates.Completed == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No update fields provided"})
		return
	}
	
	var updatedTodo Todo
	// For simplicity, we fetch the current state first.
	// A more optimized approach might build the query dynamically.
	err = db.QueryRow("SELECT id, title, completed FROM todos WHERE id = $1", id).Scan(&updatedTodo.ID, &updatedTodo.Title, &updatedTodo.Completed)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find todo"})
		}
		return
	}

	// Apply updates
	if updates.Title != nil {
		updatedTodo.Title = *updates.Title
	}
	if updates.Completed != nil {
		updatedTodo.Completed = *updates.Completed
	}
	
	_, err = db.Exec("UPDATE todos SET title = $1, completed = $2 WHERE id = $3", updatedTodo.Title, updatedTodo.Completed, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
		return
	}

	c.JSON(http.StatusOK, updatedTodo)
}

// deleteTodo handles DELETE /api/todos/:id
func deleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	result, err := db.Exec("DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check deletion status"})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
