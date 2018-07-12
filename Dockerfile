FROM golang:1.10 as build
ADD . /go/src/github.com/solcates/whereami/
WORKDIR /go/src/github.com/solcates/whereami
RUN go build


FROM alpine
COPY --from=build /go/src/github.com/solcates/whereami/whereami /whereami
ENTRYPOINT ["/whereami"]