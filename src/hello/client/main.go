package main

import (
  "flag"
  "fmt"
  "time"
  "hello/ifc"
  "v.io/v23"
  "v.io/v23/context"
  _ "v.io/x/ref/runtime/factories/generic"
)

var (
  server = flag.String(
      "server", "", "Name of the server to connect to")
)

func main() {
  ctx, shutdown := v23.Init()
  defer shutdown()
  f := ifc.HelloClient(*server)
  ctx, cancel := context.WithTimeout(ctx, time.Minute)
  defer cancel()
  hello, _ := f.Get(ctx)
  fmt.Println(hello)
}
