FROM golang:latest

COPY . .
EXPOSE 8090
RUN go build timeserver.go

CMD ["./timeserver"]
