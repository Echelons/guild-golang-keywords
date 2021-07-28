FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . ./

RUN env GOOS=linux GOARCH=amd64 go build -o /main cmd/main/main.go

CMD [ "/main" ]