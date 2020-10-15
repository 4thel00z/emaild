package debug

import "emaild/pkg/libemail"

type Debug struct{}

var (
	Module = Debug{}
)

func (Y Debug) Version() string {
	return "v1"
}

func (Y Debug) Namespace() string {
	return "debug"
}

func (Y Debug) Routes() []libemail.Route {
	return []libemail.Route{
		// Add route definitions here
		{
			Path:        "routes",
			Method:      "GET",
			CurlExample: "curl http://<addr>/<version>/<namespace>/routes",
		},
	}
}
func (Y Debug) HandlerById(i int) libemail.Service {
	switch i {
	// Add handlers for routes here
	case 0:
		return GetRoutesHandler
	}
	// This makes the server return a 404 by default
	return nil
}

func (Y Debug) LongPath(route libemail.Route) string {
	return libemail.DefaultLongPath(Y, route)
}
