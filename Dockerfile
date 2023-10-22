FROM golang:alpine

LABEL maintainer="Muhammad Avtara Khrisna"

RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

RUN mkdir /app
WORKDIR /app

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

RUN go build -o /build

EXPOSE 9000

CMD [ "/build" ]