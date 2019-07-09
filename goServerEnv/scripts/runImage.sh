#!/bin/bash

read -p "Enter val: " val

sudo docker run -e VAL="$val" -d -p 3000:8080 lesteven/goserver
