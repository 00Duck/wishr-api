#!/bin/sh
docker run -d \
-p 9191:9191 \
--name wishr-api \
--env-file ../.env \
--network host \
00duck/wishr-api:1.4.2