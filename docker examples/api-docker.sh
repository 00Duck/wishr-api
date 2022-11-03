#!/bin/sh
docker run -d \
-p 9191:9191 \
--name wishr-api \
--env-file .env \
--network wishr-net \
wishr-api:1.0