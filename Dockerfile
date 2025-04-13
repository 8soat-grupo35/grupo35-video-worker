FROM golang:1.23
RUN apt-get update && apt-get install ffmpeg -y

WORKDIR /app
COPY . /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download


CMD ["air", "-c", ".air.toml"]
