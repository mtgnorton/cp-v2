FROM --platform=linux/amd64  alpine

MAINTAINER mtgnorton

ENV WORKDIR  /app

WORKDIR $WORKDIR/

RUN apk --no-cache add tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone
ENV TZ Asia/Shanghai

COPY ./temp/linux_amd64/main $WORKDIR/main

COPY ./public $WORKDIR/public

COPY ./template $WORKDIR/template

COPY ./config $WORKDIR/config

COPY wait-for.sh $WORKDIR/wait-for

RUN chmod +x $WORKDIR/main && chmod +x $WORKDIR/wait-for


EXPOSE 8200

EXPOSE 8201

CMD ["/bin/bash","-c","-- ./main"]
