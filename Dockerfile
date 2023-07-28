FROM golang:1.20-alpine
WORKDIR /go/src/app
COPY . .
RUN go mod tidy
RUN go build -o go-server-app .
CMD ["./go-server-app"]
