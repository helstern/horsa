#!/usr/bin/env sh

while read LINE
do
    if test -z "${LINE}"
    then
        break
    else
        echo ${LINE}
    fi
done
