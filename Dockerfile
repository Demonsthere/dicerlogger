FROM golang:1.14 as builder

WORKDIR /go/src/app

ENV CGO_ENABLED 0

COPY . .

RUN go build -v -o main ./cmd/main.go
RUN mkdir /app &&\
	mv ./main /app/main

FROM alpine:3.12

WORKDIR /app

RUN apk update &&\
	apk add ca-certificates &&\
	rm -rf /var/cache/apk/*

COPY --from=builder /app /app

CMD ["/app/main"]
