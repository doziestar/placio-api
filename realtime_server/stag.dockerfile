# Use official base image of Go
ARG GOLANG_VERSION=1.21
FROM golang:${GOLANG_VERSION}-alpine3.17 AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy Go modules manifests
COPY go.mod go.sum ./
COPY pkg/go.mod pkg/go.sum ./pkg/
COPY realtime_server/go.mod realtime_server/go.sum ./realtime_server/

# Download dependencies
RUN go mod download

# Copy the entire project directory inside the container
COPY . .

ENV PORT=7080
# Build the Go app for a smaller binary size
# -v: print the names of packages as they are compiled.
# -o: name of the binary.
# -ldflags to set variable values in runtime.
# -a for "all", rebuilds everything
# -installsuffix cgo: a suffix to use in the name of the package installation directory,
# in order to keep output separate from default builds.
RUN CGO_ENABLED=0 GOOS=linux go build -v -a -installsuffix cgo -o main ./realtime_server/main.go

# Start a new stage from scratch
FROM alpine:3.17

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose port on the Docker host, so we can access it
# from the outside.
EXPOSE 7080

# Command to run the executable
CMD ["./main"]
