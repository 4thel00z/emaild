package email

import (
	"emaild/pkg/libemail"
	"emaild/pkg/libemail/filters"
	"github.com/monzo/typhon"
)

func PostSendEmailsHandler(app libemail.App) typhon.Service {
	return func(req typhon.Request) typhon.Response {
		value := req.Context.Value(filters.ValidationResult)
		email := value.(*PostSendEmailRequest)

		response := req.Response(email)
		response.StatusCode = 200
		return response
	}
}
