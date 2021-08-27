FROM alpine:3
RUN apk add --no-cache ca-certificates

COPY ./server ./

COPY ./web ./web

ENTRYPOINT ["./server"]
