FROM golang:alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go get github.com/gin-gonic/gin
RUN env GOOS=linux GOARCH=amd64 go build -o main .

EXPOSE 8080

CMD ["./main"]