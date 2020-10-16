package filters

import (
	"context"
	"emaild/pkg/libemail"
	"fmt"
	"github.com/monzo/typhon"
)
const (
	ValidationResult ="validation_result"
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

			val, err := (*route.Validator)(req)

			if err != nil {
				msg := err.Error()
				return req.Response(libemail.GenericResponse{
					Message: fmt.Sprintf("[%s] %s validation error", route.LongPath, route.Method),
					Error:   &msg,
				})
			}

			req.Context = context.WithValue(req.Context, ValidationResult, val)
			return svc(req)
		}

		return svc(req)
	}
}
