FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct

RUN sed -i 's!http://dl-cdn.alpinelinux.org/!https://mirrors.ustc.edu.cn/!g' /etc/apk/repositories && \
   apk update --no-cache && apk add --no-cache tzdata


WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
RUN go build -ldflags="-s -w" -o /app/main .



FROM --platform=linux/amd64  alpine


ENV WORKDIR  /app

WORKDIR $WORKDIR/

RUN sed -i 's!http://dl-cdn.alpinelinux.org/!https://mirrors.ustc.edu.cn/!g' /etc/apk/repositories &&  \
     apk update --no-cache && apk add --no-cache ca-certificates

COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai


COPY --from=builder  /app/main $WORKDIR/main

COPY ./public $WORKDIR/public

COPY ./template $WORKDIR/template

COPY ./config $WORKDIR/config

COPY wait-for.sh $WORKDIR/wait-for

RUN chmod +x $WORKDIR/main && chmod +x $WORKDIR/wait-for


EXPOSE 8200

EXPOSE 8201

CMD ["./main"]
