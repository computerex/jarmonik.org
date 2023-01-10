#!/bin/bash
echo "building jarmonik";
source ~/.bashrc
cd /root/jarmonik.org
rm -rf public
git pull
go run makesite.go
rm -rf /var/www/jarmonik.org/public
cp -R ./public /var/www/jarmonik.org/public
cp -R ./static /var/www/jarmonik.org/public/static