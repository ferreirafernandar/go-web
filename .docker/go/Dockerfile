#@version 3.7.5
FROM golang:1.14.6-alpine3.12 as base

WORKDIR /go/src/
COPY ./go/src ./hello_web

WORKDIR /go/src/hello_web

RUN go get -v ./...  \
    && go build -ldflags="-s -w"
RUN ls -la

CMD ["hello_web"]

#FROM scratch as prod
FROM alpine as prod

COPY --from=base /go/src/hello_web/hello_web .

CMD ["/hello_web"]