FROM golang:alpine AS builder
WORKDIR  /go/src/app 
COPY main.go .
RUN go mod init v1
RUN go mod tidy
RUN go build -o webserver .


FROM alpine 
WORKDIR /app
COPY --from=builder /go/src/app/  /app/
CMD ["./webserver"]
