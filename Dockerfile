FROM golang:1.10 as build
ADD . /go/src/github.com/solcates/whereami/
WORKDIR /go/src/github.com/solcates/whereami
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build  -i -v -o whereami


FROM alpine
COPY --from=build /go/src/github.com/solcates/whereami/whereami /whereami
ENTRYPOINT ["/whereami"]