#!/bin/bash

mkdir 0 A
touch 1 2 4 5 6 7 8 9
ln -s 0 3

chmod 401 0
chmod 402 1
chmod 604 2
chmod 510 4
chmod 460 5
chmod 460 6
chmod 510 7
chmod 604 8
chmod 402 9
chmod 401 A

TZ=UTC touch -t 198601050000 0
TZ=UTC touch -t 198611130001 1
TZ=UTC touch -t 198803050010 2
TZ=UTC touch -h -t 199002160011 3
TZ=UTC touch -t 199010070100 4
TZ=UTC touch -t 199011070101 5
TZ=UTC touch -t 199102080110 6
TZ=UTC touch -t 199103080111 7
TZ=UTC touch -t 199405201000 8
TZ=UTC touch -t 199406101001 9
TZ=UTC touch -t 199504101010 A


echo
TZ=utc ls -l --time-style='+XF XR' | sed 1d | awk '{print $1, $6, $7, $8, $9, $10 }'
tar -cf done.tar 0 1 2 3 4 5 6 7 8 9 A

