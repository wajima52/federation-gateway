FROM golang:1.19.2-alpine3.16
ENV ROOT=/go/src/app
WORKDIR ${ROOT}
RUN apk add alpine-sdk

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2

RUN go install github.com/volatiletech/sqlboiler/v4@latest
RUN go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest

CMD ["go", "run", "server.go"]