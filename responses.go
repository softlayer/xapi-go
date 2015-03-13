package xapi

type Response interface {
	Status() string
	ErrorDescription() string
}

type ResponseBase struct {
	status           string `xmlrpc:"Status"`
	errorDescription string `xmlrpc:"ErrorDescription"`
}

func (resp ResponseBase) Status() string {
	return resp.status
}

func (resp ResponseBase) ErrorDescription() string {
	return resp.errorDescription
}

type StringResponse struct {
	ResponseBase
	Value string
}

type StringsResponse struct {
	ResponseBase
	Value []string
}

type SessionResponse struct {
	ResponseBase
	Value Session
}

type VDIResponse struct {
	ResponseBase
	Value VDI
}

type VDBResponse struct {
	ResponseBase
	Value VDB
}

type VMResponse struct {
	ResponseBase
	Value VM
}

type EventResponse struct {
	ResponseBase
	Value Event
}

type EventsResponse struct {
	ResponseBase
	Value []Event
}

type VIFResponse struct {
	ResponseBase
	Value VIF
}

type PIFResponse struct {
	ResponseBase
	Value PIF
}

type HostResponse struct {
	ResponseBase
	Value Host
}
