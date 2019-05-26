# NATS Server playground
Repo contains sample code used for testing of NATS.io server concepts and embedding.

## Initial Seed list
Service looks up `_nats._tcp.nats.local` SRV record for initial seed list.
For this test it is set up with local `dnsmasq` as follows:
```
address=/nats.local/127.0.0.1
srv-host=_nats._tcp.nats.local,localhost,9190
srv-host=_nats._tcp.nats.local,localhost,9191
```

## Running
Run `run.sh <0-9>` to start services. Argument changes last digit of the port to allow synchronization (yes hacky but whatever).
Each service starts publishing to the queue every 5 seconds with its port in the message (for readability) and timestamp.

Client port is randomized (as each service only connects to itself) while cluster one is specified via `-p` flag.