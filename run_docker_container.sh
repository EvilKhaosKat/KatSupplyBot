#!/usr/bin/env bash
touch KatSupplyBot.db

#mounts local db file with container file, so that all the changes to db will be able on host machine
docker run -d --name KatSupplyBot -v $(pwd)/KatSupplyBot.db:/go/src/github.com/EvilKhaosKat/KatSupplyBot/KatSupplyBot.db katsupplybot
