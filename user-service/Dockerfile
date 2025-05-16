# Gunakan base image golang
FROM golang:1.21 AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build aplikasi
RUN go build -o main .

# Gunakan image yang lebih ringan untuk menjalankan binary
FROM alpine:latest

WORKDIR /root/

# Copy binary dari builder
COPY --from=builder /app/main .

# Expose port yang digunakan aplikasi Go (misalnya 8080)
EXPOSE 8080

# Jalankan binary
CMD ["./main"]