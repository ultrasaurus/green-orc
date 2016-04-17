

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

call server directly with vrpc:
```
$V_BIN/vrpc --v23.credentials $ORC_PATH/cred/basics call "/$IP_ADDR:$PORT" Get
```

or discover its methods:
```
$V_BIN/vrpc --v23.credentials $ORC_PATH/cred/basics signature "/$IP_ADDR:$PORT"
```
which returns the interface:
```
type "hello/ifc".Hello interface {
  // Returns a greeting.
  Get() (greeting string | error)
}
```

or just find info about a single method:
```
$V_BIN/vrpc --v23.credentials $ORC_PATH/cred/basics signature "/$IP_ADDR:$PORT" Get
```
