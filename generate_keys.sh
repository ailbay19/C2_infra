#!/bin/sh

# openssl req -new -newkey rsa:2048 -days 365 -nodes -x509 -keyout server.key -out server.crt

# openssl req -new -newkey rsa:2048 -days 365 -nodes -x509 -keyout client.key -out client.crt

openssl req -new -newkey rsa:2048 -days 365 -nodes -x509 -keyout client.key -out client.crt -addext 'subjectAltName=DNS:localhost'

openssl req -new -newkey rsa:2048 -days 365 -nodes -x509 -keyout server.key -out server.crt -addext 'subjectAltName=DNS:localhost'


rm ./vm_servers/ssl/*
rm ./nginx/ssl/*
rm ./client/ssl/*

cp server.key ./vm_servers/ssl
cp server.key ./nginx/ssl

cp client.key ./vm_servers/ssl
cp client.key ./client/ssl

cp server.crt ./vm_servers/ssl
cp server.crt ./nginx/ssl
cp server.crt ./client/ssl

cp client.crt ./vm_servers/ssl
cp client.crt ./nginx/ssl
cp client.crt ./client/ssl

rm server.key server.crt client.key client.crt