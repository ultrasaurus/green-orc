

1. install vanadium
2. clone this repo, cd into the directory

```
source .env
go install hello/server
go install hello/client
$V_BIN/principal create \
    --overwrite $ORC_PATH/cred/basics tutorial
```
