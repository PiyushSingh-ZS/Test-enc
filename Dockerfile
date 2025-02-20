# Stage 1: Build
FROM golang:1.22 AS builder

WORKDIR /app

# Copy Go modules and source files

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -buildvcs=false -o main

RUN go mod download


RUN ls -la

RUN pwd

# Stage 2: Package
FROM alpine:latest

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Give execution permissions
RUN chmod +x /main

# Expose the application port
EXPOSE 8000

# Command to run the application
CMD ["/main"]