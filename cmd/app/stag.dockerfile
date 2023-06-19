ARG GOLANG_VERSION=1.20
# Stage 1: Build the application
FROM golang:${GOLANG_VERSION}-alpine3.17 AS builder

WORKDIR /app

COPY ./go.mod .
COPY ./go.sum .
COPY ./pkg/go.mod ./pkg/
COPY ./pkg/go.sum ./pkg/
COPY ./cmd/app/go.mod ./cmd/app/
COPY ./cmd/app/go.sum ./cmd/app/

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY ./cmd/app/ ./cmd/app

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -v -a -installsuffix cgo -o app ./cmd/app/

# Stage 2: Create the runtime image
FROM alpine:3.15

# Set the working directory to /app
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/app .

# Run the binary
CMD ["./app"]
