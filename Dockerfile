FROM golang:1.24-alpine

RUN apk add --no-cache git && \
    go install github.com/air-verse/air@v1.62.0

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
EXPOSE 8080
CMD ["air"]