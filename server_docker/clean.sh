#!/bin/bash

sudo docker container rm -f $(sudo docker container ls -aq)
sudo docker image rm $(sudo docker image ls -aq)