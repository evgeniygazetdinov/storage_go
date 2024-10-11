FROM golang:alpine

ENV GO111MODULE=auto \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

  

COPY . /storage_go/


WORKDIR /storage_go


RUN go get github.com/gin-gonic/gin && go get -u github.com/go-sql-driver/mysql

EXPOSE 8080

CMD go run main.go
