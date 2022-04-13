#!/bin/bash

POSTGRES_USER=$(whoami) 
POSTGRES_PASSWORD="password" 
POSTGRES_DB="db_$(whoami)" 

if [[ ! -z $( docker version ) ]]; then
    docker run -t -p 5432:5432 --name postgres \
        -e POSTGRES_PASSWORD=$POSTGRES_PASSWORD \
        -e POSTGRES_USER=$POSTGRES_USER \
        -e POSTGRES_DB=$POSTGRES_DB  \
        -v $PWD/initdb:/docker-entrypoint-initdb.d \
        -d postgres:12 
else
    echo "Please, install docker to continue"
    exit 1
fi

