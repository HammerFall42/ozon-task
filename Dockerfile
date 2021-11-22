FROM golang:1.14.6-alpine3.12 as builder

COPY go.mod go.sum /github.com/HammerFall42/ozon-task/
WORKDIR /github.com/HammerFall42/ozon-task

RUN go mod download

COPY . /github.com/HammerFall42/ozon-task

RUN CGO_ENABLED=0 GOOS=linux go build -o ./build/app.out /github.com/HammerFall42/ozon-task/main/main.go

FROM alpine

RUN apk add --no-cache ca-certificates && update-ca-certificates


COPY --from=builder /github.com/HammerFall42/ozon-task/build/app.out .
COPY --from=builder /github.com/HammerFall42/ozon-task/config.yml .

EXPOSE 8080 8080

ENTRYPOINT ["./app.out"]