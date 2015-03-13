A XAPI ([XenServer API](http://docs.vmd.citrix.com/XenServer/6.0.0/1.0/en_gb/api/)) go binding with native interfaces and structs.  Still a work in progress, welcoming pull requests.

[Documentation](http://godoc.org/github.com/softlayer/xapi-go)

```go
import (
    "fmt"

    xapi "github.com/softlayer/xapi-go"
)

func main() {
    x := xapi.NewXapiClient("http://127.0.0.1", "root", "pass", "1.2")
    if err := x.Login(); err != nil {
        panic(err)
    }
    // obtain session object to find out the host we are connected to
    session, _ := x.GetSession()
    hostname, _ := x.GetHostname(session.ThisHost)
    host, _ := x.GetHost(session.ThisHost)
    machines, err := x.GetVMs()
    // ...
}
```

Licensed under MIT.
