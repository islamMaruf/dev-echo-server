# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o dev-echo-server .

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/dev-echo-server .

# Copy .env.example as default .env if needed
COPY --from=builder /app/.env.example ./.env.example

# Expose port
EXPOSE 3000

# Run the application
CMD ["./dev-echo-server"]
