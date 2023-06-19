ARG GOLANG_VERSION=1.20
# Stage 1: Build the application
FROM golang:${GOLANG_VERSION}-alpine3.17 AS builder

WORKDIR /app

#RUN mkdir -p /app/cmd/auth
#RUN mkdir -p /app/pkg

RUN #echo "Listing /app"
RUN #ls -la /app

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

ENV PORT=7070
# Build the application
#RUN #CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o auth ./cmd/auth

# Stage 2: Create the runtime image
#FROM alpine:3.15

# Set the working directory to /app
#WORKDIR /app

# Copy the compiled binary from the builder stage
#COPY --from=builder /app/auth .

# Copy the required files from pkg module
#COPY pkg /app/pkg
RUN go install github.com/cespare/reflex@latest

# Expose port 8080
EXPOSE 7070

# Run the binary
#CMD ["./app/auth"]
CMD ["reflex", "-r", "\\.go$", "-s", "--", "go", "run", "./cmd/app/main.go"]
