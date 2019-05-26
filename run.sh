#!/bin/sh 

CLIENT_ID=$1
PORT=909$CLIENT_ID
CLUSTER_PORT=919$CLIENT_ID

CMD="build/nats-test-linux-amd64 -p $PORT -c $CLUSTER_PORT"

echo $CMD
$CMD

