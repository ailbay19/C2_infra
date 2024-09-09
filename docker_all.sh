#!/bin/sh

sudo sh clean

sudo docker compose build --no-cache

sudo sh save_transfer