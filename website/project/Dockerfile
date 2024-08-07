FROM golang:alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go get github.com/gin-gonic/gin
RUN go build -o .

EXPOSE 8080

CMD ["./main.go"]