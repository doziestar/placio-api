ARG GOLANG_VERSION=1.20
# Stage 1: Build the application
FROM golang:${GOLANG_VERSION}-alpine3.17 AS builder

WORKDIR /app

#RUN mkdir -p /app/cmd/auth
#RUN mkdir -p /app/pkg

RUN #echo "Listing /app"
RUN #ls -la /app

COPY ./go.* .
COPY ./pkg/go.* ./pkg/
#COPY ./cmd/auth/go.* ./cmd/auth/
#COPY ./cmd/users/go.* ./cmd/users/
COPY ./cmd/app/go.* ./cmd/app/
#COPY ./cmd/business/go.* ./cmd/business/
#COPY ./cmd/chats/go.* ./cmd/chats/
#COPY ./cmd/media/go.* ./cmd/media/
#COPY ./cmd/events/go.* ./cmd/events/

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
#COPY ./pkg ./pkg
COPY ./cmd/app ./cmd/app
# COPY ./config ./config

ENV PORT=7070
# Build the application
RUN #CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/app

# Stage 2: Create the runtime image
FROM alpine:3.15

# Set the working directory to /app
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/app .

# Copy the required files from pkg module
#COPY pkg /app/pkg
# RUN go install github.com/cespare/reflex@latest

# Expose port 8080
EXPOSE 7070

# Run the binary
CMD ["./app/app"]
# CMD ["reflex", "-r", "\\.go$", "-s", "--", "go", "run", "./cmd/app/main.go"]
