# Stage 1: Build the application
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/golimiter ./cmd/server

# Stage 2: Create the final image
FROM alpine:latest

RUN apk --no-cache add ca-certificates

# Copy the built application from the builder stage
COPY --from=builder /go/bin/golimiter /go/bin/golimiter

# Expose port 8080
EXPOSE 8080

# Run the application
ENTRYPOINT ["/go/bin/golimiter"]
