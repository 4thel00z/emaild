package debug

import "emaild/pkg/libemail"

type GenericResponse struct {
	Message interface{} `json:"message"`
	Error   *string     `json:"error,omitempty"`
}

type GetRoutesResponse struct {
	Routes []libemail.Route `json:"routes"`
	Error  *string          `json:"error,omitempty"`
}
