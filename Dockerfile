FROM golang:1.22.0-alpine

# Set the working directory to /app
WORKDIR /app

# Copy the go.mod and go.sum files to the working directory
COPY go.mod go.sum /

# Download and install any required go dependencies
RUN go mod download

# Copy the entire secure code to the working directory
COPY . .

# BUild the Go application
RUN go build -o main .

# Expose the port specified by the PORT environment variable
EXPOSE 8502

# Set the entry point of the container to the executable
CMD ["./main"]