# Use an official Golang runtime as a parent image
FROM golang:1.23.1-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
