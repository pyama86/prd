FROM golang:latest
RUN apt-get update -qqy && \
    apt-get install -qqy rrdtool librrd-dev
CMD make build
