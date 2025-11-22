# Start from the latest golang base image
FROM golang:1.21-alpine AS builder

# Add Maintainer Info
LABEL maintainer="Rezky <rezky@example.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
# Assuming main.go is in the root. If it's in cmd/, change to: RUN go build -o main ./cmd/...
RUN go build -o main .

# Start a new stage from scratch
FROM alpine:latest  

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .
# Copy .env file if it exists (optional, but good practice)
# COPY --from=builder /app/.env .  

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
