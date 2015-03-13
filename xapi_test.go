package xapi_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/softlayer/xapi-go"
)

func TestLogin(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `<?xml version="1.0"?>
<methodResponse>
<params>
<param>
<value>
<struct>
<member>
<name>Status</name>
<value>Success</value>
</member>
<member>
<name>Value</name>
<value>OpaqueRef:de305d54-75b4-431b-adb2-eb6b9e546013</value>
</member>
</struct>
</value>
</param>
</params>
</methodResponse>`)
	}))
	defer ts.Close()

	x := xapi.NewClient(ts.URL, "username", "password", "1.2")
	err := x.Login()
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
}
