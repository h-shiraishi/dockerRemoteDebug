FROM golang:1.10

RUN go get github.com/labstack/echo
RUN go get github.com/labstack/echo/middleware

RUN go get github.com/derekparker/delve/cmd/dlv

COPY ./src /go/src
WORKDIR /go/src/app

RUN go build -gcflags "all=-N -l" -o /server app.go

CMD ["dlv", "--listen=:2345", "--headless=true", "--api-version=2", "exec", "/server"]