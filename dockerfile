FROM golang:1.17.5-alpine3.15
LABEL HyunSang Park <parkhyunsang@kakao.com>

RUN apk update & apk upgrade 

ENV GO111MODULE=on 

WORKDIR /app

# COPY
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY public ./
COPY . . 

RUN go build -o hyunsang-photogray

EXPOSE 80

ENTRYPOINT ["./hyunsang-photogray"]