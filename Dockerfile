FROM golang:1.14-alpine as builder
RUN apk add --no-cache upx
WORKDIR /go/src/golang-fullcycle-desafio-k8s-hpa
COPY . .
RUN GOOS=linux go build -ldflags "-s -w" -tags netgo -a main.go
RUN upx --brute main

FROM busybox:latest
WORKDIR /golang-fullcycle-desafio-k8s-hpa
COPY --from=builder /go/src/golang-fullcycle-desafio-k8s-hpa/main .
EXPOSE 8080
ENTRYPOINT ["./main"]