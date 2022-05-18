#!/bin/bash

./prepare.sh
docker-compose up -d

docker-compose ps

docker logs master
docker attach master

