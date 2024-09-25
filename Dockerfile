# Use alpine-based Go image, which is smaller and less restrictive
FROM golang:1.23-alpine

# Install necessary packages (e.g., git, and build-base for C dependencies)
RUN apk add --no-cache git build-base

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the entire project directory into the container
COPY . .

# Set the working directory to your cmd/main folder where your main.go is located
WORKDIR /app/cmd/main

# Build the Go app with verbose output
RUN go build -v -o main .

# Expose the port on which the Go app runs
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
