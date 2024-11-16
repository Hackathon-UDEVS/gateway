# Build stage
FROM golang:1.23-alpine AS builder

# Install required system packages
RUN apk add --no-cache git make protoc

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd/main.go

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/main .
# Copy any config files needed
COPY --from=builder /app/api/casbin/model.conf ./api/casbin/

# Expose the service port (from your config)
EXPOSE 6060

# Run the binary
CMD ["./main"]
