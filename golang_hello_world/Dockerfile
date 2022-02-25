FROM golang:1.16.4-alpine3.13 as build

WORKDIR /go/app

ADD . .

RUN GGO_ENABLED=0 go build -o app


FROM scratch as production

COPY --from=build /go/app/app /app

ENTRYPOINT ["/app"]