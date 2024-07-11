FROM golang:latest

# Set metadata for the image
LABEL maintainer="kananniftiyev"
LABEL version="1.0"
LABEL description="File Service"

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main ./cmd

EXPOSE 8080

CMD ["./main"]