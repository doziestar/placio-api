ARG GOLANG_VERSION=1.21
# Stage 1: Build the application
FROM golang:${GOLANG_VERSION}-alpine3.17 AS builder

WORKDIR /app

RUN #echo "Listing /app"x
RUN #ls -la /app

COPY go.mod .
COPY go.sum .
COPY ./pkg/go.mod ./pkg/
COPY ./pkg/go.sum ./pkg/
COPY ./app/go.mod ./app/
COPY ./app/go.sum ./app/


# Download dependencies
RUN go mod download
RUN go install github.com/cespare/reflex@latest

# Copy the rest of the application source code
COPY ./app/ ./app

ENV PORT=7070

EXPOSE 7070

# Run the binary
CMD ["reflex", "-r", "\\.go$", "-s", "--", "go", "run", "./app/main.go"]
