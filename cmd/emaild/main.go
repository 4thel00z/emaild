package main

import (
	"context"
	"emaild/pkg/libemail"
	"emaild/pkg/libemail/debug"
	"flag"
	"github.com/logrusorgru/aurora"
	"github.com/monzo/typhon"
	"log"
	"os"
	"os/signal"
	"strconv"
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
	port       = flag.Int("port", 1337, "<port> [defaults to 1337]")
	host       = flag.String("host", "0.0.0.0", "<host> [defaults to 0.0.0.0]")
	configPath = flag.String("config", ".emaildrc.json", "path to config. [defaults to .emaildrc.json]")
	debugFlag  = flag.Bool("debug", false, "enable replacement of <auth> with real token [defaults to false]")
	verbose    = flag.Bool("verbose", false, "enable verbose logging [defaults to false]")
)

func main() {
	flag.Parse()

	log.Println("\n", aurora.Magenta(banner), "\n")
	log.Println("👩	Version:", version)

	config, err := libemail.ParseConfig(*configPath)
	if err != nil {
		log.Fatalf("could not parse the configuration because of: %s", err.Error())
	}
	addr := *host + ":" + strconv.Itoa(*port)
	app := libemail.NewApp(addr, config, *verbose, *debugFlag, debug.Module)

	svc := app.Router.Serve().
		Filter(typhon.ErrorFilter).
		Filter(typhon.H2cFilter)
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
