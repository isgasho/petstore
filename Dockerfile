FROM golang:alpine AS build

RUN \
    apk add --update build-base make git && \
    rm -rf /var/cache/apk/*

RUN mkdir -p /src
WORKDIR /src

COPY . /src

RUN make install

FROM alpine

EXPOSE 8000/tcp

ENTRYPOINT ["/usr/local/bin/petstore"]
CMD []

COPY --from=build /go/bin/petstore /usr/local/bin/petstore

