ARG GOLANG_VERSION=1.21
# Stage 1: Build the application
FROM golang:${GOLANG_VERSION}-alpine3.17

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY ./pkg/go.mod ./pkg/
COPY ./pkg/go.sum ./pkg/
COPY ./app/go.mod ./app/
COPY ./app/go.sum ./app/


# Download dependencies
RUN go mod download
RUN go install github.com/cespare/reflex@latest
RUN #go install github.com/cosmtrek/air@latest

# Copy the rest of the application source code
COPY . .

ENV PORT=7070

EXPOSE 7070

# Run the binary
CMD ["reflex", "-r", "\\.go$", "-s", "--", "go", "run", "./app/main.go"]
