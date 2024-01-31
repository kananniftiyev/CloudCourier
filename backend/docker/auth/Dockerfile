FROM golang:latest

# Set metadata for the image
LABEL maintainer="kananniftiyev"
LABEL version="1.0"
LABEL description="Auth Service"

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY cmd/auth cmd/auth
COPY internal/auth internal/auth
COPY utils utils
COPY docs docs
COPY .env .env
COPY firebase-x.json firebase.json


RUN go build -o main ./cmd/auth

EXPOSE 8080

CMD ["./main"]
