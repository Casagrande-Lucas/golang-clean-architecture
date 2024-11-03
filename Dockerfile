# Stage 1: Build the Go application
FROM golang:1.22-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o golang-clean-architecture ./cmd/server

# Stage 2: Create a lightweight image for running
FROM alpine:latest

# Install bash (if needed)
RUN apk add --no-cache bash

# Set the working directory
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/golang-clean-architecture .

# Copy the .env.example file
COPY .env.example .env

# Copy wait-for-it.sh script
COPY wait-for-it.sh /app/wait-for-it.sh
RUN chmod +x /app/wait-for-it.sh

# Expose port 8080
EXPOSE 8080

# Command to run the binary
CMD ["./wait-for-it.sh", "db:3306", "--", "./golang-clean-architecture"]
