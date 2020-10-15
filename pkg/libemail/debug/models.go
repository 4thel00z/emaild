package debug

import "emaild/pkg/libemail"

type GetRoutesResponse struct {
	Routes []libemail.Route `json:"routes"`
	Error  *string          `json:"error,omitempty"`
}

type PostValidatorTestRequest struct {
	HasToBeThere string `json:"has_to_be_there" validate:"empty=false"`
}
