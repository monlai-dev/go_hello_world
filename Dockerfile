FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

# Copy the source code
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/app

# Final stage
FROM alpine:3.19

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Expose port
EXPOSE 3000

# Command to run the application
CMD ["./main"]