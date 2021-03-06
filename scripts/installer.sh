#!/usr/bin/env bash
PATH=/bin:/sbin:/usr/bin:/usr/sbin:/usr/local/bin:/usr/local/sbin:~/bin
export PATH
#=================================================================#
#   System Required:  Now Ubuntu only                             #
#   Description: One click config oj env                          #
#   Author: miloas <genesis.null@gmail.com>                       #
#==================================================================

clear
echo
echo "#############################################################"
echo "# One click config oj env, lol!                             #"
echo "# Author: miloas <genesis.null@gmail.com>                   #"
echo "# Github: https://github.com/Aetema/oj                      #"
echo "#############################################################"
echo

#Current folder
cur_dir=`pwd`

# Make sure only root can run our script
rootness(){
    if [[ $EUID -ne 0 ]]; then
       echo "Error:This script must be run as root!" 1>&2
       exit 1
    fi
}

get_char(){
    SAVEDSTTY=`stty -g`
    stty -echo
    stty cbreak
    dd if=/dev/tty bs=1 count=1 2> /dev/null
    stty -raw
    stty echo
    stty $SAVEDSTTY
}

# Pre-installation settings
pre_install(){
    echo
    echo "Press any key to start...or Press Ctrl+C to cancel"
    char=`get_char`
    #Install necessary dependencies
    apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 0C49F3730359A14518585931BC711F9BA15703C6
    echo "deb [ arch=amd64,arm64 ] http://repo.mongodb.org/apt/ubuntu xenial/mongodb-org/3.4 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-3.4.list
    apt-get -y update
    apt-get install -y wget unzip gzip curl make libseccomp2 mongodb-org g++ default-jre default-jdk
    echo
}

download_runtime(){
    cd ${cur_dir}
    if ! wget http://download.redis.io/releases/redis-3.2.8.tar.gz; then
        echo "Failed to download redis-3.2.8"
        exit 1
    fi
    tar xzf redis-3.2.8.tar.gz
    if [ $? -eq 0 ]; then
        echo "Decompress redis-3.2.8.tar.gz success"
    else
        echo "Decompress redis-3.2.8.tar.gz failed"
        exit 1
    fi
    cd redis-3.2.8
    make
    cd src
    make install 
    cd ../utils
    ./install_server.sh
    service redis_6379 start

    cd ${cur_dir}
    rm redis-3.2.8.tar.gz

    git clone https://github.com/quark-zju/ljudge
    cd ljudge
    make 
    make install 
    cp etc/ljudge /etc/ljudge -r
}


config_env(){
    rootness
    pre_install
    download_runtime
}

config_env


clear
echo
echo "#############################################################"
echo "#                       快去改变世界                          #"
echo "#############################################################"
echo
