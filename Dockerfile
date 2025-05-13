# Stage 1: Build the Go application
FROM golang:1.24 AS builder

# Set the working directory in the container
WORKDIR /app

# Copy the go.mod and go.sum files separately to take advantage of caching
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the application
RUN CGO_ENABLED=0 go build -o ping-server .

# Stage 2: Use a minimal Alpine image as the base
FROM alpine:3.21

# Install BusyBox
RUN apk add --no-cache busybox curl

# Set the working directory in the container
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/ping-server .

# Expose the port the server will run on
EXPOSE 8080

# Add labels to the Docker image
LABEL org.label-schema.name="Ping Server"
LABEL org.label-schema.description="A simple ping server that responds with pong."
LABEL org.label-schema.version=${{ github.ref_name }}
LABEL org.label-schema.vcs-url="https://github.com/homelab-peej/ping-server"
LABEL org.label-schema.vcs-ref=${{ github.sha }}

# Run the application
CMD ["./ping-server"]
