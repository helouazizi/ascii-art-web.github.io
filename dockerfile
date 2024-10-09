# Use an official Golang image as a base
FROM golang:alpine

# Set the working directory to /app
WORKDIR /ascii


# Copy the application code
COPY . .

# Build the Go application
RUN go build -o main main.go

# Expose the port the application will use
EXPOSE 8080

# Run the command to start the application when the container starts
CMD ["go", "run", "main.go"]