FROM golang:1.14-alpine as builder

WORKDIR /gotodo/
ADD . ./
RUN go build -o main ./src/

FROM alpine:latest

WORKDIR /app/
COPY --from=builder /gotodo/main ./

RUN chmod +x main

CMD ["./main"]