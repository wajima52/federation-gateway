FROM golang:1.19.2-alpine3.16
WORKDIR /var/www/

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2

CMD ["go", "run", "server.go"]