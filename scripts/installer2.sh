#!/usr/bin/env bash
PATH=/bin:/sbin:/usr/bin:/usr/sbin:/usr/local/bin:/usr/local/sbin:~/bin

if ! wget https://github.com/quark-zju/lrun/releases/download/v1.1.4/lrun_1.1.4_amd64.deb; then
    echo "Failed to download lrun"
    exit 1
fi  
sudo dpkg -i lrun_1.1.4_amd64.deb
sudo gpasswd -a $USER lrun

rm lrun_1.1.4_amd64.deb

service mongod start


if ! wget https://storage.googleapis.com/golang/go1.8.linux-amd64.tar.gz; then
    echo "Failed to download golang1.8"
    exit 1
fi
tar -zxvf go1.8.linux-amd64.tar.gz -C /usr/local/
rm go1.8.linux-amd64.tar.gz
 
echo 'export GOROOT=/usr/local/go' >> ~/.bashrc
echo 'export GOPATH=$HOME/go' >> ~/.bashrc
echo 'export PATH=$PATH:$GOROOT/bin:$GOPATH/bin' >> ~/.bashrc
source ~/.bashrc
    
go get github.com/codegangsta/negroni
go get github.com/Miloas/oj/middleware