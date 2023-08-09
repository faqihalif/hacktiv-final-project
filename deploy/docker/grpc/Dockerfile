FROM golang:alpine as builder

WORKDIR /app

COPY . .

RUN go build -o app cmd/grpc/main.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/app .

CMD [ "./app" ]