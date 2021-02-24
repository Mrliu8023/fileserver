FROM golang:1.16 AS build

WORKDIR /app/gt

COPY main.go .

COPY embedgo .

COPY files files

ENV GO111MODULE off

RUN ./embedgo build -v -o server

FROM busybox:glibc

COPY --from=build /app/gt/server server

ENV WEBPORT=8000

CMD ["./server"]
