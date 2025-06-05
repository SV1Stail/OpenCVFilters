#!/usr/bin/env bash

if [ $# -eq 0 ]; then
    echo "Usage: $0 [options]"
    echo " -r       : docker compose up --build"
    echo " -f       : docker compose down -v && docker compose build --no-cache && docker compose up"
    echo " -d       : docker compose down -v"
    echo " -s       : docker compose up"
    echo " -h       : this msg"
    echo " NO flags : show this help"
    exit 0
fi

while getopts "hrfds" OPTS; do
case $OPTS in
h)
echo " -r       : docker compose up --build"
echo " -f       : docker compose down -v && docker compose build --no-cache && docker compose up"
echo " -d       : docker compose down -v"
echo " -s       : docker compose up"
echo " -h       : this msg"
echo " NO flags : help"
;;
r)
docker compose up --build
;;
f)
# docker compose down -v && docker compose up --build --force-recreate
docker compose down -v && docker compose build --no-cache && docker compose up
;;
d)
docker compose down -v
;;
s)
docker compose up
;;
\?)
echo " -r       : docker compose up --build"
echo " -f       : docker compose down -v && docker compose build --no-cache && docker compose up"
echo " -d       : docker compose down -v"
echo " -s       : docker compose up"
echo " NO flags : docker compose up --build"
;;
esac
done