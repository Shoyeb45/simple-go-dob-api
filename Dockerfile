FROM golang:1.25.5-alpine3.23

RUN go version
# Install build dependencies
RUN apk add --no-cache git ca-certificates tzdata

# set up the work directory
WORKDIR /app

# Copy only go.mod and go.sum for better caching
COPY go.mod go.sum ./

# Download the dependenices
RUN go mod download

# Copy the source code 
COPY . .

# Build the application, -ldflags for smaller binary size
RUN go build -ldflags='-w -s -extldflags "-static"' -o main ./cmd/server

EXPOSE 8080

CMD ["./main"]