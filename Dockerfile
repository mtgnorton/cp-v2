FROM --platform=linux/amd64  alpine

MAINTAINER mtgnorton


ENV WORKDIR  /app

WORKDIR $WORKDIR/

COPY ./temp/linux_amd64/main $WORKDIR/main

COPY ./public $WORKDIR/public

COPY ./template $WORKDIR/template

COPY ./config $WORKDIR/config

COPY wait-for.sh $WORKDIR/wait-for

RUN chmod +x $WORKDIR/main && chmod +x $WORKDIR/wait-for


EXPOSE 8200

EXPOSE 8201

CMD ["/bin/bash","-c","-- ./main"]
