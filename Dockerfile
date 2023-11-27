# Use an official Go runtime as a base image
FROM golang:latest AS builder

WORKDIR /app/gorest

COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the local package files to the container
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Use a smaller base image for the final stage
FROM alpine:latest

WORKDIR /app/gorest

# Copy the compiled binary from the builder stage
COPY --from=builder /app/gorest/main .

# Expose port 5000 to the outside world
EXPOSE 5000

# Command to run the executable
CMD ["./main"]
