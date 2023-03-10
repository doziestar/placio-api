ARG GOLANG_VERSION=1.20.1
FROM golang:${GOLANG_VERSION}-alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Start fresh from a smaller image
FROM alpine:latest

RUN apk --no-cache add ca-certificates

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main /app/main

# Expose port 8080 to the outside world
EXPOSE 7071

# Command to run the executable
CMD ["/app/main"]