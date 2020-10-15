package senders

import (
	"emaild/pkg/libemail"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"io/ioutil"
	"os"
	"testing"
)

func TestGmailSenderInit(t *testing.T) {
	credsPath, found := os.LookupEnv("GMAIL_CREDENTIALS_PATH")
	if !found {
		t.Fatal("GMAIL_CREDENTIALS_PATH env was not set!")
	}

	tokenPath, found := os.LookupEnv("GMAIL_TOKEN_PATH")
	if !found {
		t.Fatal("GMAIL_TOKEN_CONFIG_PATH env was not set!")
	}

	creds, err := ioutil.ReadFile(credsPath)
	if err != nil {
		t.Fatal(err)
	}
	tokenConfig, err := google.ConfigFromJSON(creds, gmail.GmailSendScope)
	if err != nil {
		t.Fatal(err)
	}
	file, err := os.Open(tokenPath)
	if err != nil {
		t.Fatal(err)
	}
	token, err := libemail.TokenFromReader(file)
	if err != nil {
		t.Fatal(err)
	}
	g := &GmailSender{}
	err = g.Init(tokenConfig, token)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGmailSenderSend(t *testing.T) {
	credsPath, found := os.LookupEnv("GMAIL_CREDENTIALS_PATH")
	if !found {
		t.Fatal("GMAIL_CREDENTIALS_PATH env was not set!")
	}

	tokenPath, found := os.LookupEnv("GMAIL_TOKEN_PATH")
	if !found {
		t.Fatal("GMAIL_TOKEN_CONFIG_PATH env was not set!")
	}

	gmailTo, found := os.LookupEnv("GMAIL_TO")
	if !found {
		t.Fatal("GMAIL_TO env (which is who will receive the email send out by the test) was not set!")
	}

	creds, err := ioutil.ReadFile(credsPath)
	if err != nil {
		t.Fatal(err)
	}
	tokenConfig, err := google.ConfigFromJSON(creds, gmail.GmailSendScope)
	if err != nil {
		t.Fatal(err)
	}
	file, err := os.Open(tokenPath)
	if err != nil {
		t.Fatal(err)
	}
	token, err := libemail.TokenFromReader(file)
	if err != nil {
		t.Fatal(err)
	}
	g := &GmailSender{}
	err = g.Init(tokenConfig, token)
	if err != nil {
		t.Fatal(err)
	}
	body := libemail.Base64("VGVzdCBFbWFpbAo=")
	message := &libemail.Email{
		To:      []string{gmailTo},
		Subject: "Emaild Test email",
		Body:    &body,
	}
	err = g.Send(message)
	if err != nil {
		t.Fatal(err)
	}
}
