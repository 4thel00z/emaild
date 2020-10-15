package debug

import (
	"emaild/pkg/libemail"
	"github.com/monzo/typhon"
)

func GetRoutesHandler(app libemail.App) typhon.Service {
	return func(req typhon.Request) typhon.Response {

		response := req.Response(&GetRoutesResponse{
			Routes: app.Routes(),
		})

		response.StatusCode = 200
		return response
	}
}
