FROM golang:1.14.0

LABEL author=scottlin

COPY Makefile /app

RUN make install && make compile

RUN make exec
