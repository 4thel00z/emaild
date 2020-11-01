package senders

import (
	"emaild/pkg/libemail"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

func TestGmailSenderInit(t *testing.T) {

	credsPath, found := os.LookupEnv("GMAIL_CREDENTIALS_PATH")
	if !found {
		t.Skip("GMAIL_CREDENTIALS_PATH env was not set!")
	}

	tokenPath, found := os.LookupEnv("GMAIL_TOKEN_PATH")
	if !found {
		t.Skip("GMAIL_TOKEN_CONFIG_PATH env was not set!")
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

func TestGmailSenderInitCI(t *testing.T) {

	creds, found := os.LookupEnv("GMAIL_CREDENTIALS")
	if !found {
		t.Skip("GMAIL_CREDENTIALS env was not set!")
	}

	tokenRaw, found := os.LookupEnv("GMAIL_TOKEN")
	if !found {
		t.Skip("GMAIL_TOKEN_CONFIG env was not set!")
	}

	tokenConfig, err := google.ConfigFromJSON([]byte(creds), gmail.GmailSendScope)
	if err != nil {
		t.Fatal(err)
	}

	token, err := libemail.TokenFromReader(strings.NewReader(tokenRaw))
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
		t.Skip("GMAIL_CREDENTIALS_PATH env was not set!")
	}

	tokenPath, found := os.LookupEnv("GMAIL_TOKEN_PATH")
	if !found {
		t.Skip("GMAIL_TOKEN_CONFIG_PATH env was not set!")
	}

	gmailTo, found := os.LookupEnv("GMAIL_TO")
	if !found {
		t.Fatal("GMAIL_TO env (which is who will receive the gmail send out by the test) was not set!")
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

	body := "Hello from emaild!"
	message := &libemail.Email{
		To:      []string{gmailTo},
		Subject: "Emaild Test gmail",
		Body:    &body,
	}
	response, err := g.Send(message)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("successfully sent: %v\n", response)

}

func TestGmailSenderSendCI(t *testing.T) {
	creds, found := os.LookupEnv("GMAIL_CREDENTIALS")
	if !found {
		t.Skip("GMAIL_CREDENTIALS env was not set!")
	}

	tokenRaw, found := os.LookupEnv("GMAIL_TOKEN")
	if !found {
		t.Skip("GMAIL_TOKEN_CONFIG env was not set!")
	}

	tokenConfig, err := google.ConfigFromJSON([]byte(creds), gmail.GmailSendScope)
	if err != nil {
		t.Fatal(err)
	}

	token, err := libemail.TokenFromReader(strings.NewReader(tokenRaw))
	if err != nil {
		t.Fatal(err)
	}

	gmailTo, found := os.LookupEnv("GMAIL_TO")
	if !found {
		t.Fatal("GMAIL_TO env (which is who will receive the gmail send out by the test) was not set!")
	}

	g := &GmailSender{}
	err = g.Init(tokenConfig, token)
	if err != nil {
		t.Fatal(err)
	}

	body := "Hello from emaild!"
	message := &libemail.Email{
		To:      []string{gmailTo},
		Subject: "Emaild Test gmail",
		Body:    &body,
	}
	response, err := g.Send(message)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("successfully sent: %v\n", response)

}
