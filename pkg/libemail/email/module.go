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
			CurlExample: "curl -X POST -d @examples/send_email.json  http://0.0.0.0:1337/v1/email/send",
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
