FROM golang:alpine as builder

RUN apk update && apk upgrade && apk add --update make
COPY . /root/cncf
RUN cd /root/cncf && ls -la && make build

FROM alpine
RUN apk update && apk upgrade
WORKDIR /root/
COPY --from=builder /root/cncf/bin/http .
EXPOSE 80
ENV HTTP_PORT=80
CMD ["./http"] 

