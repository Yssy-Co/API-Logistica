FROM alpine:latest
RUN apk add git go curl
ADD main.go .
RUN go get gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer
RUN go get github.com/sirupsen/logrus
RUN go build main.go
RUN chmod +x main
CMD ["./main"]