#!/bin/bash

docker run -d --name app --network c2bridge app
docker run -d --name nginx --network c2bridge -p 443:443 -p 80:80 nginx