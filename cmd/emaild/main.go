package main

import (
	"context"
	"emaild/pkg/libemail"
	"emaild/pkg/libemail/debug"
	"emaild/pkg/libemail/filters"
	internalGmail "emaild/pkg/libemail/gmail"
	"flag"
	"github.com/logrusorgru/aurora"
	"github.com/monzo/typhon"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

const (
	banner = `
                                   /$$ /$$       /$$
                                  |__/| $$      | $$
  /$$$$$$  /$$$$$$/$$$$   /$$$$$$  /$$| $$  /$$$$$$$
 /$$__  $$| $$_  $$_  $$ |____  $$| $$| $$ /$$__  $$
| $$$$$$$$| $$ \ $$ \ $$  /$$$$$$$| $$| $$| $$  | $$
| $$_____/| $$ | $$ | $$ /$$__  $$| $$| $$| $$  | $$
|  $$$$$$$| $$ | $$ | $$|  $$$$$$$| $$| $$|  $$$$$$$
 \_______/|__/ |__/ |__/ \_______/|__/|__/ \_______/`
	version = "0.0.1"
)

var (
	port                  = flag.Int("port", 1337, "<port> [defaults to 1337]")
	host                  = flag.String("host", "0.0.0.0", "<host> [defaults to 0.0.0.0]")
	configPath            = flag.String("config", ".emaildrc.json", "path to config. [defaults to .emaildrc.json]")
	debugFlag             = flag.Bool("debug", false, "enable replacement of <auth> with real token [defaults to false]")
	verbose               = flag.Bool("verbose", false, "enable verbose logging [defaults to false]")
	modules               = flag.String("modules", internalGmail.Namespace, "comma separated list of modules to load")
	googleOauthConfigPath = flag.String("google-oauth-config", "credentials.json", "path to google oauth config file. [defaults to credentials.json]")
	googleOauthTokenPath  = flag.String("google-oauth-token", "token.json", "path to google oauth token file. [defaults to token.json]")
)

func main() {
	flag.Parse()

	log.Println("\n", aurora.Magenta(banner))
	log.Println("\n👩	Version:", version)

	config, err := libemail.ParseConfig(*configPath)
	if err != nil {
		log.Fatalf("could not parse the configuration because of: %s", err.Error())
	}
	addr := *host + ":" + strconv.Itoa(*port)

	toLoad := []libemail.Module{debug.Module}
	for _, moduleName := range strings.Split(*modules, ",") {

		switch strings.ToLower(moduleName) {
		case internalGmail.Namespace:

			creds, err := ioutil.ReadFile(*googleOauthConfigPath)
			if err != nil {
				log.Fatal(err)
			}
			config, err := google.ConfigFromJSON(creds, gmail.GmailSendScope)
			if err != nil {
				log.Fatalf("Unable to parse google client secret file to config: %e", err)
			}

			token, err := libemail.LoadToken(*googleOauthTokenPath)
			if err != nil {
				log.Fatalf("Unable to parse google auth token file: %e", err)
			}

			module, err := internalGmail.NewModule(internalGmail.WithDebug(*debugFlag), internalGmail.WithTokenConfig(config, token))
			if err != nil {
				log.Fatal(err)
			}
			toLoad = append(toLoad, module)
		}
	}

	app := libemail.NewApp(addr, config, *verbose, *debugFlag, toLoad...)

	svc := app.Router.Serve().
		Filter(typhon.ErrorFilter).
		Filter(typhon.H2cFilter).
		Filter(filters.Validation(app))
	srv, err := typhon.Listen(svc, addr)
	if err != nil {
		panic(err)
	}

	log.Printf("🏁	Listening on %v", srv.Listener().Addr())
	if app.Debug {
		app.PrintConfig()
	}
	app.PrintRoutes(srv.Listener().Addr().String())
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
	log.Printf("☠️  Shutting down in max 10 sec..")
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	srv.Stop(c)
}
