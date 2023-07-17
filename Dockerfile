FROM golang:latest

WORKDIR /app

# Copy the entire root directory into the container
COPY ./ /app

# Run go run main.go inside the /app directory
CMD ["go", "run", "main.go"]
