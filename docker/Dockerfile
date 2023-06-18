FROM golang:alpine
WORKDIR /build
COPY ./ ./
EXPOSE 8080
RUN go build -o main server.go
ENTRYPOINT ["/build/main"]


