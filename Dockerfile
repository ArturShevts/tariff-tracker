FROM golang:1.22-alpine

WORKDIR /app

# Install dependencies
RUN apk add --no-cache git

# Copy go mod files first for better layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the application
RUN go build -o server ./cmd/server/main.go

# Run the application
CMD ["./server"]