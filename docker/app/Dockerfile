# Build stage
FROM golang:1.23.4-alpine AS builder

# Install git and ca-certificates
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download && go mod verify

# Copy source code
COPY . .

# Build the main application binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo -o main .

# Build the migration binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo -o migrate ./cmd/migrate.go

# Final stage - minimal alpine for production
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk update && apk add --no-cache ca-certificates && \
    rm -rf /var/cache/apk/*

# Create appuser
ENV USER=appuser
ENV UID=10001

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/home/appuser" \
    --shell "/sbin/nologin" \
    --uid "${UID}" \
    "${USER}"

# Create necessary directories with proper permissions
RUN mkdir -p /home/appuser && \
    chown -R appuser:appuser /home/appuser

# Set working directory
WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/main ./main
COPY --from=builder /app/migrate ./migrate

# Ensure the binaries are executable
RUN chmod +x ./main ./migrate

# Change ownership of app directory to appuser
RUN chown -R appuser:appuser /app

# Use the unprivileged user
USER appuser:appuser

# Expose port
EXPOSE 8080

# Default command (can be overridden)
CMD ["./main"]