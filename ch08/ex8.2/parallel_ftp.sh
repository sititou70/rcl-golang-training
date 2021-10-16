#!/bin/sh
set -eux

HOST="localhost"
PORT="2121"
FTP_COMMAND="
user name pass
get main.go
pwd
cd sample
ls
binary
get image.jpg
quit"

ftpJob() {
  i=$1

  WORK_DIR="TEMP_${i}"
  [ ! -e $WORK_DIR ] && mkdir $WORK_DIR
  cd $WORK_DIR

  echo "$FTP_COMMAND" | ftp -n $HOST $PORT
}

for i in $(seq 1000); do
  ftpJob $i &
done

wait
