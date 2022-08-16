FROM golang:latest

# NOTE: the resulting image is very large, we should use a multi-stage build here

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -a ./server/main.go

EXPOSE 8080
EXPOSE 8081

CMD "./main"
