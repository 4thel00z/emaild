package debug

import (
	"emaild/pkg/libemail"
)

type Debug struct{}

var (
	Module = Debug{}
)

const (
	Namespace = "debug"
)

func (d Debug) Version() string {
	return "v1"
}

func (d Debug) Namespace() string {
	return Namespace
}

func (d Debug) Routes() []libemail.Route {
	return []libemail.Route{
		// Add route definitions here
		{
			Path:        "routes",
			Method:      "GET",
			CurlExample: "curl http://<addr>/<version>/<namespace>/routes",
		},
		{
			Path:        "validator_test",
			Method:      "POST",
			CurlExample: "curl -X POST --data '{\"has_to_be_there\":\"something\"}' http://<addr>/<version>/<namespace>/validator_test",
			Validator:   libemail.GenerateJSONValidator(PostValidatorTestRequest{}),
		},
	}
}
func (d Debug) HandlerById(i int) libemail.Service {
	switch i {
	// Add handlers for routes here
	case 0:
		return GetRoutesHandler
	}
	// This makes the server return a 404 by default
	return nil
}

func (d Debug) LongPath(route libemail.Route) string {
	return libemail.DefaultLongPath(d, route)
}
