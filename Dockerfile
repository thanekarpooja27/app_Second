# Use the official Golang image as the base image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code to the container
COPY . .

# Build the Go application
RUN go build -o main

# Expose the port your application will run on
EXPOSE 8080

# Command to run the application
CMD ["./main"]
