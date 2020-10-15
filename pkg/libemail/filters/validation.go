package filters

import (
	"emaild/pkg/libemail"
	"fmt"
	"github.com/monzo/typhon"
)

func Validation(app libemail.App) typhon.Filter {
	return func(req typhon.Request, svc typhon.Service) typhon.Response {
		pattern := app.Router.Pattern(req)
		for _, route := range app.Routes() {
			if route.LongPath != pattern {
				continue
			}

			if route.Validator == nil {
				return svc(req)
			}

			err := (*route.Validator)(req)

			if err != nil {
				msg := err.Error()
				return req.Response(libemail.GenericResponse{
					Message: fmt.Sprintf("[%s] %s validation error", route.LongPath, route.Method),
					Error:   &msg,
				})
			}
		}

		return svc(req)
	}
}
