FROM golang:1.23

WORKDIR /app
COPY . /app

RUN apt update && apt install -y ffmpeg
RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download


CMD ["air", "-c", ".air.toml"]
