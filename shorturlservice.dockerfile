FROM golang:1.16-alpine

WORKDIR /shorturlservice/


COPY . ./

RUN ls -la ./*

RUN go mod download -x

ENTRYPOINT ["sh", "./build.sh"]