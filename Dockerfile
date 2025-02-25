# Build stage
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main .

# Final stage
FROM alpine:3.21
RUN addgroup -S basic && adduser -S basic -G basic
WORKDIR /app
COPY --from=builder --chown=basic:basic /app/main .
USER basic
EXPOSE 5000
CMD ["./main"]