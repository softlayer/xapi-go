package xapi

type ResponseBase struct {
	Status           string
	ErrorDescription string
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
