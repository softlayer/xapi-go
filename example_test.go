// Example methods
package xapi_test

import (
	"xapi"
)

func ExampleRpcCall() {
	x := xapi.NewXapiClient("http://localhost/", "username", "password", "1.2")
	host := xapi.Host{}
	err := x.RpcCall(&host, "host.get_record", "324c2264-d86f-4a42-a971-bb5fd6203877")
	if err != nil {
		fmt.Printf("%v", host)
	}
}
