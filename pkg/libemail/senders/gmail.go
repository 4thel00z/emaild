package senders

import (
	"context"
	"emaild/pkg/libemail"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/oauth2"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
	"log"
	"strings"
)

type GmailSender struct {
	Service *gmail.Service
}

func (g *GmailSender) Init(config *oauth2.Config, token *oauth2.Token) error {
	log.Printf("GmailSender.Init(%v,%v)\n", config, token)
	service, err := gmail.NewService(context.Background(), option.WithHTTPClient(config.Client(context.Background(), token)))
	if err != nil {
		return err
	}
	g.Service = service
	return nil
}

func (g *GmailSender) Cleanup() error {
	log.Println("GmailSender.Cleanup()")
	return nil
}

func (g *GmailSender) Send(message *libemail.Email) error {
	header := make(map[string]string)
	header["From"] = message.From
	header["To"] = strings.Join(message.To, ";")
	header["Cc"] = strings.Join(message.Cc, ";")
	header["Bcc"] = strings.Join(message.Bcc, ";")
	header["Subject"] = message.Subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	var payload string

	if message.Body != nil {
		unpacked, err := message.Body.Unpack()
		if err != nil {
			return err
		}
		payload = unpacked
	} else if message.HTML != nil {
		unpacked, err := message.HTML.Unpack()
		if err != nil {
			return err
		}
		payload = unpacked

		header["Content-Type"] = "text/html; charset=\"utf-8\""
	} else {
		return errors.New("message.Body or message.HTML must be set. message.File not supported as of now")
	}

	var msg string
	for k, v := range header {
		msg += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	msg += "\r\n" + payload

	response, err := g.Service.Users.Messages.Send("me", &gmail.Message{
		Raw: base64.RawURLEncoding.EncodeToString([]byte(msg)),
	}).Do()
	if err != nil {
		return err
	}
	fmt.Printf("successfully sent: %v\n", response)
	return nil
}
