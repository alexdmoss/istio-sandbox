FROM golang:1.14.13-alpine3.12 as builder

RUN mkdir /build 
WORKDIR /build 

ADD ./ /build/

RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

RUN adduser -u 1000 -D -g '' appuser

RUN go get -d -v ./...

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go test -v -ldflags="-w -s" -c -o main.test

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main .

# --------------------------------------------

FROM scratch 

WORKDIR /app

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /build/ /app/

USER appuser

EXPOSE 8080

ENTRYPOINT ["./main"]
