FROM golang:1.16.3 as builder
WORKDIR /go/src/petProject
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/app/main.go



FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/petProject/app .

#EXPOSE 8080
ENTRYPOINT ["./app"]