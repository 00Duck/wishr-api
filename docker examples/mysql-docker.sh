#!/bin/sh
docker run -d -p 3306:3306 \
-e MYSQL_DATABASE=wishr \
-e MYSQL_USER=wishr \
-e MYSQL_PASSWORD=wishr \
-e MYSQL_ROOT_PASSWORD="your super strong root pw" \
--name wishr-db \
--network wishr-net \
-v /srv/wishr/db:/var/lib/mysql \
mysql:8.0.31