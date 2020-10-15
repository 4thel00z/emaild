package main

import (
	"emaild/pkg/libemail"
	"flag"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"io/ioutil"
	"log"
)

var (
	credsPath = flag.String("config", "credentials.json", "path to token config [defaults to credentials.json]")
	outPath   = flag.String("out", "token.json", "where to write the token to [defaults to token.json]")
)

func main() {
	flag.Parse()
	creds, err := ioutil.ReadFile(*credsPath)
	if err != nil {
		log.Fatal(err)
	}

	config, err := google.ConfigFromJSON(creds, gmail.GmailSendScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %e", err)
	}

	token, err := libemail.GetTokenFromWeb(config)
	if err != nil {
		log.Fatalf("libemail.GetTokenFromWeb(%v): %e", config, err)
	}
	err = libemail.SaveToken(*outPath, token)
	if err != nil {
		log.Fatalf("libemail.SaveToken(%v,%v): %e", outPath, token, err)
	}
}
