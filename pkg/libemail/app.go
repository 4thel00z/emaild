package libemail

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/monzo/typhon"
	"log"
	"strings"
)

type App struct {
	Addr    string   `json:"addr"`
	Config  Config   `json:"config"`
	Modules []Module `json:"modules"`
	Router  *typhon.Router
	Debug   bool
	Verbose bool
}

func NewApp(addr string, config Config, verbose, debug bool, modules ...Module) App {

	app := App{
		Addr:    addr,
		Config:  config,
		Modules: modules,
		Debug:   debug,
		Verbose: verbose,
	}

	router := &typhon.Router{}

	for _, module := range modules {
		for i, route := range module.Routes() {
			path := module.LongPath(route)
			handler := module.HandlerById(i)
			if handler == nil {
				handler = Default404Handler
			}
			router.Register(strings.ToUpper(route.Method), path, handler(app))
		}
	}

	app.Router = router

	return app
}

func (app App) Routes() []Route {
	var routes []Route
	addr := app.Addr

	for _, module := range app.Modules {
		version := module.Version()
		namespace := module.Namespace()

		for _, route := range module.Routes() {
			route.CurlExample = strings.ReplaceAll(route.CurlExample, "<addr>", addr)
			route.CurlExample = strings.ReplaceAll(route.CurlExample, "<version>", version)
			route.CurlExample = strings.ReplaceAll(route.CurlExample, "<namespace>", namespace)
			route.CurlExample = strings.ReplaceAll(route.CurlExample, "<path>", route.Path)
			if app.Debug {
				// Add module wise injections of f.e. the <auth> tag
			}
			route.LongPath = module.LongPath(route)
			routes = append(routes, route)

		}
	}
	return routes
}
func (app App) PrintRoutes(addr string) {
	routes := app.Routes()
	if len(routes) > 0 {
		log.Println("👠\tThe routes 🛣️  are:")
	}
	for _, route := range routes {
		log.Printf("\thttp://%v%s with method: %s", addr, route.LongPath, route.Method)
		log.Printf("\tQuery this endpoint like this:\n\t\t%s", route.CurlExample)

	}
}

func (app App) Register(module Module) {
	for i, route := range module.Routes() {
		path := module.LongPath(route)
		handler := module.HandlerById(i)
		fmt.Println("HANDLER", handler, handler(app))
		if handler == nil {
			handler = Default404Handler
		}
		fmt.Println("METHOD", route.Method, "PATH:", path)
		app.Router.Register(strings.ToUpper(route.Method), path, handler(app))
	}

}

func (app App) PrintConfig() {
	fmt.Print("======================\t✈️\tConfig incoming\t✈️\t======================\n")
	spew.Dump(app.Config)
	fmt.Print("======================\t✈️ Config landed! ✈️\t======================\n")
}
