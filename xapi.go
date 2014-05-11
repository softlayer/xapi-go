// A generic binding of the Citrix XenServer API.  Read more here: http://docs.vmd.citrix.com/XenServer/6.0.0/1.0/en_gb/api/
package xapi

import (
	"fmt"
	"github.com/kolo/xmlrpc"
	"net"
	"net/http"
	"time"
)

type XapiClient struct {
	Session  string
	Uri      string
	Username string
	Password string
	Version  string
	rpc      *xmlrpc.Client
}

type APIResult struct {
	Status string
	Value interface{}
	ErrorDescription string
}

// Stand up a new XapiClient.  Version should probably be "1.2" unless you know what you are doing.
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

// Check to see if a field that should be an OpaqueRef is actually empty.  Sometimes it's an empty string (rare)
// but other times it's OpaqueRef:NULL which is good to know.
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

// Get hostname of a Host.  Useful in combination with GetSession and session.This_host
func (client *XapiClient) GetHostname(opref string) (hostname string, err error) {
	err = client.SessionCall(&hostname, "host.get_hostname", opref)
	return
}

// Useful for making multiple calls that require the session ID.  Automatically prepends the existing
// session OpaqRef to the beginning of the API call.  You can see the session ID by looking at
// XapiClient.Session.
func (client *XapiClient) SessionCall(result interface{}, call string, params ...interface{}) (err error) {
	if client.Session == "" {
		return fmt.Errorf("NO_SESSION")
	}
	p := make([]interface{}, len(params)+1)
	p[0] = client.Session
	err = client.RpcCall(result, call, append(p, params...)...)
	return
}

// Custom Dialer for HTTP so that the initial connection only lasts for 1 minute
// and that the lifetime of the connection is only 1 minute as well. See http://golang.org/pkg/net/#Conn
// You shouldn't need to use this directly.
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

// Make a generic RPC call passing in a pointer to a struct (or f). The call parameter
// is a combination of class.message.  For example: VIF.get_record, host.evacuate, pool.eject.
// Any time the XAPI specifies a `type ref` it's really an OpaqueReference, which is a UUID, and
// as far as xmlrpc and like library are concerned, a string.
//		x := xapi.NewXapiClient("http://localhost/", "username", "password", "1.2")
//		host := xapi.Host{}
//		err := x.RpcCall(&host, "host.get_record", "324c2264-d86f-4a42-a971-bb5fd6203877")
//		if err != nil {
//			fmt.Printf("%v", host)
//		}
func (client *XapiClient) RpcCall(result interface{}, call string, params ...interface{}) (error) {
	response := APIResult{Value: result}

	err := client.rpc.Call(call, params, &response)
	if err != nil {
		return err
	}

	if response.Status != "Success" {
		return fmt.Errorf("XenServer Failed: %s", response.ErrorDescription)
	}

	return nil
}
