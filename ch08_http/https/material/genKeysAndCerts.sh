#!/bin/bash

# set up server ca and client ca
# private keys
openssl genrsa -out serverCA.key 2048
openssl genrsa -out clientCA.key 2048

# cert sign requests
openssl req -new -key serverCA.key -out serverCA.csr -config serverCA.conf
openssl req -new -key clientCA.key -out clientCA.csr -config clientCA.conf

# get self-signed ca certs
openssl x509 -req -days 365 -in serverCA.csr -signkey serverCA.key -out serverCA.pem
openssl x509 -req -days 365 -in clientCA.csr -signkey clientCA.key -out clientCA.pem

# get private keys and signed certs of server and client
# private keys
openssl genrsa -out server.key 2048
openssl genrsa -out client.key 2048

# cert sign requests
openssl req -new -key server.key -out server.csr -config localhost.conf
openssl req -new -key client.key -out client.csr -config localhost.conf

# get signed certs from ca
openssl x509 -req -days 365 -CA serverCA.pem -CAkey serverCA.key -CAcreateserial -in server.csr -out server.pem
openssl x509 -req -days 365 -CA clientCA.pem -CAkey clientCA.key -CAcreateserial -in client.csr -out client.pem
