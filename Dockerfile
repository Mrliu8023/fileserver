FROM golang:latest AS build

WORKDIR /app/embedweb

COPY main.go .

RUN go build -v -o server

FROM busybox:glibc

COPY --from=build /app/embedweb/server server

ENV WEBPORT=8000

CMD ["./server"]
