# Use the official Golang image as the base image
FROM golang:1.20-alpine AS build

# Set the working directory to /app
WORKDIR /app

# Copy go.mod and go.sum to the working directory
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code to the working directory
COPY . .

# Build the Go application
RUN go build -o dating_app cmd/main.go

# Use a minimal Alpine image for the final image
FROM alpine:latest

# Set the working directory to /app
WORKDIR /app

# Copy the binary from the build stage to the final image
COPY --from=build /app/dating_app .

# Expose the port the app runs on
EXPOSE 8080

CMD ["./dating_app"]

