#!/bin/sh

openssl genrsa -des3 -passout pass:x -out service.pass.key 2048
openssl rsa -passin pass:x -in service.pass.key -out service.key
openssl req -new -key service.key -out service.csr -subj "/CN=daneharrigan.com/O=./C=US"
openssl x509 -req -days 365 -in service.csr -signkey service.key -out service.crt
openssl x509 -in service.crt -out service.pem -outform PEM
rm service.pass.key service.csr service.crt

go build -o bin/geordi-server server/*.go
#go build -o bin/geordi-cli cli/*.go
