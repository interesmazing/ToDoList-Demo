# Stage 1: Build the application
# Use the official Golang image as a builder.
FROM golang:1.21-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod file first to leverage Docker cache
COPY go.mod ./

# Download all dependencies. Dependencies will be cached if the go.mod file is not changed.
# This command will also generate the go.sum file inside the container.
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
# CGO_ENABLED=0 is important for creating a static binary that doesn't rely on C libraries.
# This makes the resulting image smaller and more portable.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Stage 2: Create the final, minimal image
# Use a minimal alpine image for a small footprint
FROM alpine:latest

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
