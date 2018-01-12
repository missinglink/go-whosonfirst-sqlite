# https://blog.docker.com/2016/09/docker-golang/
# https://blog.golang.org/docker

# docker build -t wof-sqlite-index .
# docker run -it -e REPOS='whosonfirst-data' wof-sqlite-index

FROM golang:alpine AS build-env

# https://github.com/gliderlabs/docker-alpine/issues/24

RUN apk add --update alpine-sdk

ADD . /go-whosonfirst-sqlite

RUN cd /go-whosonfirst-sqlite; make bin

FROM alpine

RUN apk add --update bzip2 curl git

# SOMETHING SOMETHING SOMETHING git lfs
# https://askubuntu.com/questions/799341/how-to-install-git-lfs-on-ubuntu-16-04

# SOMETHING SOMETHING SOMETHING awscli

VOLUME /usr/local/data

WORKDIR /go-whosonfirst-sqlite/bin/

COPY --from=build-env /go-whosonfirst-sqlite/bin/wof-sqlite-index /bin/wof-pip-server
COPY --from=build-env /go-whosonfirst-sqlite/docker/entrypoint.sh /bin/entrypoint.sh

EXPOSE 8080

ENTRYPOINT bin/entrypoint.sh

