package gmail

import (
	"emaild/pkg/libemail"
	"emaild/pkg/libemail/filters"
	"github.com/monzo/typhon"
)

func PostSendEmailsHandler(app libemail.App) typhon.Service {

	return func(req typhon.Request) typhon.Response {
		value := req.Context.Value(filters.ValidationResult)
		email := value.(*PostSendGmailRequest)
		module := app.Modules[Namespace]
		gmailModule := module.(*Gmail)
		gmailResponse, err := gmailModule.Send(email)
		if err != nil {
			tmp := err.Error()
			response := req.Response(libemail.GenericResponse{
				Message: "",
				Error:   &tmp,
			})
			response.StatusCode = 503
			return response
		}
		response := req.Response(gmailResponse)
		response.StatusCode = 200
		return response
	}
}
