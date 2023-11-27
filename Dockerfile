FROM golang:1.21-alpine3.18
WORKDIR /app
COPY app/ .
RUN go build -o helloworld .
CMD ["/app/helloworld"]
