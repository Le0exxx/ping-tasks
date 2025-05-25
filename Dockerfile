# Stage 1: Build the Go binary
FROM golang:1.24-bullseye AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod (for better caching)
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go binary
RUN go build -o server .

# Stage 2: Run the Go binary in a small image
FROM debian:bullseye-slim

# Install certificates (required for HTTPS or external calls)
RUN apt-get update && \
    apt-get install -y ca-certificates && \
    apt-get install -y libc6 && \
    rm -rf /var/lib/apt/lists/*

# Set the working directory
WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/server .

# Expose the port your web service uses (change as needed)
EXPOSE 8080

# Run the binary
CMD ["./server"]