A XAPI go binding with native interfaces and structs.  Still a work in progress, welcoming pull requests.

```go
import (
    "fmt"
    xapi "github.com/softlayer/xapi-go"
)

func main() {
    x := xapi.NewXapiClient("http://127.0.0.1", "root", "pass", "1.2")
    if err := x.Login(); if err != nil {
        panic(err)
    }
    // obtain session object to find out the host we are connected to
    session, _ := x.GetSession()
    hostname, _ := x.GetHostname(session.This_host)
    host, _ := x.GetHost(session.This_host)
    machines, err := x.GetVMs()
    // ...
}
```

Licensed under MIT.
