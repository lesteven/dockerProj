#!/bin/bash

read -p "Enter val: " val

sudo docker run -e VAL="$val" lesteven/gohello
