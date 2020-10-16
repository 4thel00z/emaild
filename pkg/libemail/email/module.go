package email

import (
	"emaild/pkg/libemail"
)

type Email struct{}

var (
	Module = Email{}
)

func (e Email) Version() string {
	return "v1"
}

func (e Email) Namespace() string {
	return "email"
}


func (e Email) Routes() []libemail.Route {
	return []libemail.Route{

		{
			Path:        "send",
			Method:      "POST",
			//TODO: add example in examples folder
			CurlExample: "curl -X POST http://<addr>/<version>/<namespace>/send",
			Validator:   libemail.GenerateJSONValidator(PostSendEmailRequest{}),
		},
	}
}
func (e Email) HandlerById(i int) libemail.Service {
	switch i {
	// Add handlers for routes here
	case 0:
		return PostSendEmailsHandler
	}
	// This makes the server return a 404 by default
	return nil
}

func (e Email) LongPath(route libemail.Route) string {
	return libemail.DefaultLongPath(e, route)
}
