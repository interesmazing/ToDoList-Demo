# Stage 1: Build the Go application
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy all source code first. This is less cache-efficient but more robust.
COPY . .

# Run go mod tidy to ensure go.mod and go.sum are in sync with the code.
RUN go mod tidy

# Build the application
# CGO_ENABLED=0 is important for a static binary.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Stage 2: Create a minimal final image
FROM alpine:latest

WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Expose port 8081 for the backend service
EXPOSE 8081

# Command to run the executable
CMD ["./main"]
