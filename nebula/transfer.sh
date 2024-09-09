#!/bin/bash

if [ "$#" -ne 1 ]; then
  echo "Usage: $0 <lighthouse_ip>"
  exit 1
fi

# Read the lighthouse IP from the argument
lighthouse_ip="$1"

scp -p 1345 -r manager user@localhost:/home/user
scp -p 1346 -r worker1 user@localhost:/home/user
scp -p 1347 -r worker2 user@localhost:/home/user

scp -i ~/test_ec2_1.pem -r lighthouse ubuntu@${lighthouse_ip}:/home/ubuntu
