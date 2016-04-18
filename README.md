

1. install vanadium
2. clone this repo, cd into the directory

```
source .env
go install hello/server
go install hello/client
$V_BIN/principal create \
    --overwrite cred/orc orc
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
$V_BIN/vrpc --v23.credentials $ORC_PATH/cred/sarah call "/$JUDY_SERVER" Get
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


## Manually transfer a blessing

Given two Principals (Judy with "judy" blessing and Sarah with the "sarah" blessing), Judy blesses Sarah as her friend. This lets Sarah's client access Judy's server.

0. Judy makes sure her firewall settings allow incoming connections, and starts her server: ```bin/server --v23.credentials cred/judy --v23.tcp.address ":$PORT" &```
1. Sarah: ```export PUBLIC_KEY=`$V_BIN/principal --v23.credentials cred/sarah get publickey` ```
   then sends $PUBLIC_KEY to Judy (via text message)
2. Judy: ```export BLESSING=`echo $PUBLIC_KEY | $V_BIN/principal bless --v23.credentials cred/judy --for=24h - friend:sarah` ```
   then sends $BLESSING to Sarah (via text message)
3. Sarah: ```echo $BLESSING | $V_BIN/principal --v23.credentials $V_TUT/cred/sarah set forpeer - judy```
4. Sarah:

```
$V_BIN/principal --v23.credentials cred/sarah dump

Public key : 20:70:2f:5f:36:c0:f8:b5:f4:bf:13:0d:83:91:70:38
Default Blessings : sarah
---------------- BlessingStore ----------------
Default Blessings                sarah
Peer pattern                     Blessings
...                              sarah
judy                             judy:friend:sarah
---------------- BlessingRoots ----------------
Public key                                        Pattern
21:7d:83:ad:9b:8b:25:a3:b6:e9:1e:d9:19:25:73:21   [judy]
20:70:2f:5f:36:c0:f8:b5:f4:bf:13:0d:83:91:70:38   [sarah]

```
