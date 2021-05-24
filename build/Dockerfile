FROM golang:1.16-alpine as builder

WORKDIR /app
ADD . ./

RUN ls -al && go build -o main main.go

FROM scratch as runtime

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /app
COPY --from=builder /app/main main

CMD [ "/app/main" ]
