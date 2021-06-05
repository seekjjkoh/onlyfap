FROM alpine:3
RUN apk add --no-cache ca-certificates

COPY ./server ./

ENTRYPOINT ["./server"]
