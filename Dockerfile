# Use the official Go image as the base image
FROM golang:1.21 AS builder

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy the local package files to the container's workspace
COPY . .

# Build the Go app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -v -o app

# Use a minimal Alpine image for the final build
FROM alpine:3.14

# Copy the binary from the builder stage to the current stage
COPY --from=builder /go/src/app/app /app

# Expose port 50051 for the gRPC server to listen on
EXPOSE 50051

# Set the binary as the entrypoint of the container
ENTRYPOINT ["/app", "--action=serve"]
