# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Run unit tests before building the binary
RUN go test -v ./unit_test/... 

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/app

# Final stage
FROM alpine:3.19

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/main .

# Expose port
EXPOSE 3000

# Use exec form for CMD
CMD ["/bin/sh", "-c", "echo \"PORT=$PORT\" > .env && \
    echo \"RENDER_DATABASE_URL=$RENDER_DATABASE_URL\" >> .env && \
    echo \"DATABASE_URL=$DATABASE_URL\" >> .env && \
    echo \"JWT_SECRET=$JWT_SECRET\" >> .env && \
    echo \"STRIPE_SECRET=$STRIPE_SECRET\" >> .env && \
    echo \"CLIENT_ID=$CLIENT_ID\" >> .env && \
    echo \"API_KEY=$API_KEY\" >> .env && \
    echo \"CHECK_SUM_KEY=$CHECK_SUM_KEY\" >> .env && \
    echo \"RENDER_REDIS_URL=$RENDER_REDIS_URL\" >> .env && \
    echo \"REDIS_URL=$REDIS_URL\" >> .env && \
    echo \"EMAIL_USERNAME=$EMAIL_USERNAME\" >> .env && \
    echo \"EMAIL_PASSWORD=$EMAIL_PASSWORD\" >> .env && \
    echo \"RABBITMQ_URL=$RABBITMQ_URL\" >> .env && \
    echo \"ENV=$ENV\" >> .env && \
    exec ./main"]
