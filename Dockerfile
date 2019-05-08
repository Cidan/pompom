FROM golang:1.12-alpine
ENV GO111MODULE on
WORKDIR /go/src/app
COPY . .
RUN apk add git && go build ./... && go install ./...
FROM alpine:3.9.3
COPY --from=0 /go/bin/pompom /pompom
ENTRYPOINT [ "/pompom" ]