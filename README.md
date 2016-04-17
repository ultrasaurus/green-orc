

1. install vanadium
2. clone this repo, cd into the directory

```
source .env
go install hello/server
go install hello/client
$V_BIN/principal create \
    --overwrite cred/basics tutorial
```

```
export IP_ADDR=ifconfig | sed -En 's/127.0.0.1//;s/.*inet (addr:)?(([0-9]*\.){3}[0-9]*).*/\2/p'
export PORT=23000
```

run the server (in the background)
```
bin/server     --v23.credentials cred/basics     --v23.tcp.address :$PORT &

```

run the client
```
bin/client  --v23.credentials cred/basics     --server /$IP_ADDR:$PORT
```
