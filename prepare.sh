#!/bin/bash

root_dir=$(pwd)
echo $root_dir

# this is unnecesary unles you changed sth in proto
#cd $root_dir/host
#make proto
#
#cd $root_dir/master
#make proto

cd $root_dir/images
docker build -f master.Dockerfile -t playground/master ..
docker build -f host.Dockerfile -t playground/host ..

cd $root_dir/config
python3 make_configs.py

cd $root_dir
cp config/docker-compose.yaml .
