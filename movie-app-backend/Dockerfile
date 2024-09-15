# Dockerfile for GoFiber backend
FROM golang:1.20-alpine

WORKDIR /app

# Copy go.mod and go.sum to download dependencies first
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the Go application
RUN go build -o movie-app-backend

# Expose the port the application runs on
EXPOSE 3000

# Command to run the application
CMD ["./movie-app-backend"]
