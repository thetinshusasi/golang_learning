# Start from the latest golang base image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

EXPOSE 3000

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY ./src/ ./src/

# # Build the Go app
RUN go build -o ./appname  ./src/main.go 


# # Expose port 3000 to the outside world

# # Command to run the executable
CMD ["./appname"]
