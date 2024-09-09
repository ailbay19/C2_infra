#!/bin/bash

rm ./server_docker/*.tar

docker save -o ./server_docker/app.tar app
docker save -o ./server_docker/nginx.tar nginx

scp -P 1345 -r server_docker user@localhost:/home/user
