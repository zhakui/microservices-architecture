FROM golang:1.12

WORKDIR /go/src/app

ADD product-service product-service
RUN chmod 777 product-service

EXPOSE 2536
EXPOSE 8080

CMD ["./product-service"]


