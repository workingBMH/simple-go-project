FROM golang:1.18

# Install Docker client
RUN apt-get update && apt-get install -y docker.io

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Command to run the executable
CMD ["./main"]
