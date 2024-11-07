#! /bin/bash
# nginxをdocker runする
docker run -d --name nginx nginx
# nginxのイメージハッシュをAPP_IDに保存
APP_ID=$(docker inspect nginx --format='{{.Id}}')
# PID namespaceを共有するコンテナを起動(https://hub.docker.com/r/brendanburns/topz/tags)
docker run --pid=container:${APP_ID} \
    -p 8080:8080 \
    brendanburns/topz:db0fa58 \
    /server -addr=0.0.0.0:8080 # topzサーバーを0.0.0.0:8080でリッスン
