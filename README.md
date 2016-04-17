

1. install vanadium
2. clone this repo, cd into the directory

```
source .env
go install hello/server
go install hello/client
$V_BIN/principal create \
    --overwrite cred/basics tutorial
```

run the server (in the background)
```
bin/server --v23.credentials $ORC_PATH/cred/basics --v23.tcp.address ":$PORT" &

```

run the client
```
bin/client --v23.credentials $ORC_PATH/cred/basics --server "/$IP_ADDR:$PORT"
```
