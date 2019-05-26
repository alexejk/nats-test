#!/bin/sh 

CLIENT_ID=$1
PORT=919$CLIENT_ID

CMD="build/nats-test-linux-amd64 -p $PORT"

echo $CMD
$CMD

