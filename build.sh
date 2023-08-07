#!/bin/bash
docker container rm -f gincms
docker image rm -f gincms
docker image build --no-cache -t gincms -f Dockerfile .
docker container run -itd --name=gincms --restart=unless-stopped -p 8066:8066 -v /usr/local/docker/gincms/dockerData/log:/www/log gincms
echo '完成!'