<template>
  <div id="app-container">
    <header>
      <h1>ToDoList Demo</h1>
      <p class="subtitle">Built with Go, Vue.js, PostgreSQL & Docker</p>
    </header>
    <main>
      <div class="todo-input-container">
        <input
          v-model="newTodoTitle"
          @keyup.enter="addTodo"
          placeholder="What needs to be done?"
          class="todo-input"
        />
        <button @click="addTodo" class="add-btn">Add</button>
      </div>

      <div v-if="loading" class="loading-state">Loading...</div>
      <div v-if="error" class="error-state">{{ error }}</div>

      <ul v-if="!loading && !error" class="todo-list">
        <li
          v-for="todo in todos"
          :key="todo.id"
          :class="{ completed: todo.completed }"
          class="todo-item"
        >
          <div class="todo-content">
            <input
              type="checkbox"
              :checked="todo.completed"
              @change="toggleTodoStatus(todo)"
              class="todo-checkbox"
            />
            <span class="todo-title">{{ todo.title }}</span>
          </div>
          <button @click="deleteTodo(todo.id)" class="delete-btn">×</button>
        </li>
      </ul>
       <div v-if="!loading && todos.length === 0" class="empty-state">
        No tasks yet. Add one above!
      </div>
    </main>
    <footer>
      <p>Created by GitHub Copilot</p>
    </footer>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';

// The base URL for the API. In a Docker Compose setup, the browser accesses the backend
// via the host machine's port, not the internal container network. The request is then
// proxied by Nginx or the Vite dev server to the backend container.
const API_URL = '/api';

const todos = ref([]);
const newTodoTitle = ref('');
const loading = ref(true);
const error = ref(null);

// Fetch all todos from the backend
const fetchTodos = async () => {
  try {
    loading.value = true;
    error.value = null;
    const response = await axios.get(`${API_URL}/todos`);
    todos.value = response.data || [];
  } catch (err) {
    console.error('Error fetching todos:', err);
    error.value = 'Failed to load tasks. Please try again later.';
  } finally {
    loading.value = false;
  }
};

// Add a new todo
const addTodo = async () => {
  if (!newTodoTitle.value.trim()) {
    return;
  }
  try {
    const response = await axios.post(`${API_URL}/todos`, {
      title: newTodoTitle.value,
    });
    todos.value.unshift(response.data); // Add to the top of the list
    newTodoTitle.value = '';
  } catch (err) {
    console.error('Error adding todo:', err);
    error.value = 'Failed to add task.';
  }
};

// Delete a todo
const deleteTodo = async (id) => {
  try {
    await axios.delete(`${API_URL}/todos/${id}`);
    todos.value = todos.value.filter((todo) => todo.id !== id);
  } catch (err) {
    console.error('Error deleting todo:', err);
    error.value = 'Failed to delete task.';
  }
};

// Toggle the completion status of a todo
const toggleTodoStatus = async (todoToUpdate) => {
  try {
    const updatedTodo = { ...todoToUpdate, completed: !todoToUpdate.completed };
    const response = await axios.put(`${API_URL}/todos/${todoToUpdate.id}`, {
      completed: updatedTodo.completed,
    });
    // Find the index and update the todo in the array to reflect the change
    const index = todos.value.findIndex((todo) => todo.id === todoToUpdate.id);
    if (index !== -1) {
      todos.value[index] = response.data;
    }
  } catch (err) {
    console.error('Error updating todo:', err);
    error.value = 'Failed to update task status.';
  }
};

// Fetch todos when the component is mounted
onMounted(fetchTodos);
</script>

<style>
:root {
  --background-color: #f4f7f6;
  --container-bg: #ffffff;
  --text-color: #333;
  --primary-color: #4a90e2;
  --primary-hover: #357abd;
  --danger-color: #e24a4a;
  --danger-hover: #c93e3e;
  --border-color: #e0e0e0;
  --completed-color: #999;
  --shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

body {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif;
  background-color: var(--background-color);
  color: var(--text-color);
  margin: 0;
  padding: 20px;
  display: flex;
  justify-content: center;
  align-items: flex-start;
  min-height: 100vh;
}

#app-container {
  width: 100%;
  max-width: 600px;
  background-color: var(--container-bg);
  border-radius: 8px;
  box-shadow: var(--shadow);
  padding: 2rem;
  box-sizing: border-box;
}

header {
  text-align: center;
  margin-bottom: 2rem;
  border-bottom: 1px solid var(--border-color);
  padding-bottom: 1rem;
}

h1 {
  color: var(--primary-color);
  margin: 0;
}

.subtitle {
  color: #777;
  font-size: 0.9rem;
  margin-top: 0.5rem;
}

.todo-input-container {
  display: flex;
  gap: 10px;
  margin-bottom: 1.5rem;
}

.todo-input {
  flex-grow: 1;
  padding: 12px 15px;
  border: 1px solid var(--border-color);
  border-radius: 6px;
  font-size: 1rem;
  transition: border-color 0.2s;
}

.todo-input:focus {
  outline: none;
  border-color: var(--primary-color);
}

.add-btn {
  padding: 12px 20px;
  background-color: var(--primary-color);
  color: white;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-size: 1rem;
  font-weight: 500;
  transition: background-color 0.2s;
}

.add-btn:hover {
  background-color: var(--primary-hover);
}

.loading-state, .error-state, .empty-state {
  text-align: center;
  padding: 2rem;
  color: #777;
  background-color: #fafafa;
  border-radius: 6px;
}

.error-state {
  color: var(--danger-color);
  background-color: #fff0f0;
}

.todo-list {
  list-style: none;
  padding: 0;
  margin: 0;
}

.todo-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 5px;
  border-bottom: 1px solid var(--border-color);
  transition: background-color 0.2s;
}

.todo-item:last-child {
  border-bottom: none;
}

.todo-item:hover {
  background-color: #f9f9f9;
}

.todo-content {
  display: flex;
  align-items: center;
  gap: 12px;
}

.todo-checkbox {
  width: 20px;
  height: 20px;
  cursor: pointer;
}

.todo-title {
  font-size: 1.1rem;
  transition: color 0.2s;
}

.todo-item.completed .todo-title {
  text-decoration: line-through;
  color: var(--completed-color);
}

.delete-btn {
  background: none;
  border: none;
  color: #aaa;
  font-size: 1.5rem;
  cursor: pointer;
  padding: 5px 10px;
  border-radius: 50%;
  line-height: 1;
  transition: color 0.2s, background-color 0.2s;
}

.delete-btn:hover {
  color: var(--danger-color);
  background-color: #fdeeee;
}

footer {
  text-align: center;
  margin-top: 2rem;
  padding-top: 1rem;
  border-top: 1px solid var(--border-color);
  font-size: 0.8rem;
  color: #aaa;
}
</style>
