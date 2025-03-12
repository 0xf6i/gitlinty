FROM golang:1.23.4-alpine

# Install dependencies
RUN apk add --no-cache git gcc musl-dev curl tar

# Install gitleaks
RUN curl -sSfL https://github.com/gitleaks/gitleaks/releases/download/v8.18.1/gitleaks_8.18.1_linux_x64.tar.gz | tar -xz -C /usr/local/bin

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main

# Default command
CMD ["./main"]