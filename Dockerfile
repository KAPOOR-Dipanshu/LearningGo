# Build stage
FROM golang:1.24-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Copy source code
COPY src/ ./src/

# Download dependencies
RUN go mod tidy
RUN go mod download

# Build the binary with optimizations
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./src

# Runtime stage
FROM alpine:3.20

# Install ca-certificates for HTTPS requests (if needed)
RUN apk --no-cache add ca-certificates

# Create a non-root user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Set working directory
WORKDIR /app

# Copy the binary from builder stage
COPY --from=builder /app/main .

# Change ownership to non-root user
RUN chown appuser:appgroup main

# Switch to non-root user
USER appuser

# Expose port (assuming default 8080, but configurable via env)
EXPOSE 8080

# Run the binary
CMD ["./main"]