FROM golang
MAINTAINER EvilKhaosKat <evilkhaoskat@gmail.com>

RUN go get github.com/EvilKhaosKat/KatSupplyBot
WORKDIR /go/src/github.com/EvilKhaosKat/KatSupplyBot

ADD token admins ./

ENTRYPOINT ./launch.sh