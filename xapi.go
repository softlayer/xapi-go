package xenserver

import (
	"fmt"
	"github.com/kolo/xmlrpc"
	"net"
	"net/http"
	"time"
	"unicode"
	"unicode/utf8"
)

/* Needed
 *  session.logout
 *  VBD.get_record
 *  VDI.get_record
 *  network.get_record
 */

/* Completed
 *  PIF.get_record
 *  VM.get_all
 *  VM.get_record
 *  VIF.get_record
 *  host.get_record
 *  host.get_hostname
 *  session.get_record
 *  session.login_with_password
 *  event.register event.unregister event.next
 */

type XapiClient struct {
	Session  string
	Uri      string
	Username string
	Password string
	Version  string
	rpc      *xmlrpc.Client
}

func NewXapiClient(uri, username, password, version string) (client XapiClient) {
	client.Uri = uri
	client.Username = username
	client.Password = password
	client.Version = version
	client.rpc, _ = xmlrpc.NewClient(
		client.Uri,
		&http.Transport{
			Dial: TimeoutDialer(),
		})

	return
}

func OpaqueRefIsEmpty(a string) bool {
	if a == "OpaqueRef:NULL" || a == "" {
		return true
	}

	return false
}

func (client *XapiClient) Login() (err error) {
	err = client.RpcCall(
		&client.Session,
		"session.login_with_password",
		client.Username,
		client.Password,
		client.Version)
	return
}

func (client *XapiClient) GetSession() (session Session, err error) {
	err = client.SessionCall(&session, "session.get_record",
		client.Session)
	return

}

func (client *XapiClient) RegisterEvent(classes ...interface{}) error {
	return client.SessionCall(nil, "event.register", classes)
}

func (client *XapiClient) UnregisterEvent() error {
	var classes []interface{}
	classes = make([]interface{}, 1)
	classes[0] = "*"
	return client.SessionCall(nil, "event.unregister", classes)
}

func (client *XapiClient) NextEvent() (ev []Event, err error) {
	err = client.SessionCall(&ev, "event.next")
	return
}

func (client *XapiClient) GetVMs() (machines VirtualMachines, err error) {
	err = client.SessionCall(&machines, "VM.get_all")
	return
}

func (client *XapiClient) GetVM(opref string) (vm VM, err error) {
	err = client.SessionCall(&vm, "VM.get_record", opref)
	return
}

func (client *XapiClient) GetVIF(opref string) (vif VIF, err error) {
	err = client.SessionCall(&vif, "VIF.get_record", opref)
	return
}

func (client *XapiClient) GetPIF(opref string) (pif PIF, err error) {
	err = client.SessionCall(&pif, "PIF.get_record", opref)
	return
}

func (client *XapiClient) GetHost(opref string) (host Host, err error) {
	err = client.SessionCall(&host, "host.get_record", opref)
	return
}

func (client *XapiClient) GetHostname(opref string) (hostname string, err error) {
	err = client.SessionCall(&hostname, "host.get_hostname", opref)
	return
}

func (client *XapiClient) SessionCall(result interface{}, call string, params ...interface{}) (err error) {
	if client.Session == "" {
		return fmt.Errorf("NO_SESSION")
	}
	p := make([]interface{}, len(params)+1)
	p[0] = client.Session
	err = client.RpcCall(result, call, append(p, params...)...)
	return
}

func TimeoutDialer() func(net, addr string) (c net.Conn, err error) {
	return func(netw, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(netw, addr, time.Minute)
		if err != nil {
			return nil, err
		}
		conn.SetDeadline(time.Now().Add(time.Minute))
		return conn, nil
	}
}

func (client *XapiClient) RpcCall(result interface{}, call string, params ...interface{}) (err error) {
	response := xmlrpc.Struct{}
	p := xmlrpc.Params{}
	p.Params = make([]interface{}, len(params))
	for i, d := range params {
		p.Params[i] = d
	}

	err = client.rpc.Call(call, p, &response)

	if err != nil {
		return
	}

	if err = checkResponse(response); err != nil {
		return
	}

	if result != nil {
		unMarshallXmlRPC(response, result)
	}

	return
}

func checkResponse(res xmlrpc.Struct) error {
	var success interface{}
	var ok bool
	var error_string interface{}

	if success, ok = res["Status"]; ok && success == "Success" {
		if _, ok = res["Value"]; !ok {
			return fmt.Errorf("'Value' is missing in result!")
		}
		return nil
	}

	if !ok {
		return fmt.Errorf("'Status' not in result!")
	}

	if error_string, ok = res["ErrorDescription"]; !ok {
		error_string = "Missing Error description!"
	}

	return fmt.Errorf("XenServer Failed: %s", error_string)
}

func UF(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[n:]
}
