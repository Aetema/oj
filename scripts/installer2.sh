#!/usr/bin/env bash
PATH=/bin:/sbin:/usr/bin:/usr/sbin:/usr/local/bin:/usr/local/sbin:~/bin

if ! wget https://github.com/quark-zju/lrun/releases/download/v1.1.4/lrun_1.1.4_amd64.deb; then
    echo "Failed to download lrun"
    exit 1
fi  
sudo dpkg -i lrun_1.1.4_amd64.deb
sudo gpasswd -a $USER lrun

rm lrun_1.1.4_amd64.deb

sudo service mongod start
