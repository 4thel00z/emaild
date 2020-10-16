package gmail

import (
	"emaild/pkg/libemail"
	"emaild/pkg/libemail/senders"
	"errors"
	"golang.org/x/oauth2"
)

type Gmail struct {
	Debug  bool
	Sender *senders.GmailSender
}

func (g Gmail) Send(message *PostSendGmailRequest) (interface{}, error) {
	email := libemail.Email{
		Account:     message.Account,
		To:          message.To,
		From:        message.From,
		Cc:          message.Cc,
		Bcc:         message.Bcc,
		Subject:     message.Subject,
		ReplyTo:     message.ReplyTo,
		Sender:      message.Sender,
		Attachments: message.Attachments,
		Body:        message.Body,
		HTML:        message.HTML,
		File:        message.File,
		Delay:       message.Delay,
	}

	return g.Sender.Send(&email)
}

type Option func(e *Gmail) error

func WithDebug(debug bool) Option {
	return func(g *Gmail) error {
		g.Debug = debug
		return nil
	}
}

func WithTokenConfig(config *oauth2.Config, token *oauth2.Token) Option {
	return func(g *Gmail) error {
		sender := &senders.GmailSender{Debug: g.Debug}
		err := sender.Init(config, token)
		if err != nil {
			return err
		}
		g.Sender = sender
		return nil
	}
}

func WithGmailSender(sender *senders.GmailSender) Option {
	return func(g *Gmail) error {
		g.Sender = sender
		sender.Debug = g.Debug
		return nil
	}
}

func NewModule(opts ...Option) (*Gmail, error) {
	g := &Gmail{}
	for _, opt := range opts {
		// Call the option giving the instantiated
		// *House as the argument
		err := opt(g)
		if err != nil {
			return nil, err
		}
	}
	if g.Sender == nil {
		return nil, errors.New("the gmail module cannot have a nil sender")
	}

	return g, nil
}

const (
	Namespace = "gmail"
)

func (e Gmail) Version() string {
	return "v1"
}

func (e Gmail) Namespace() string {
	return Namespace
}

func (e Gmail) Routes() []libemail.Route {
	return []libemail.Route{
		{
			Path:        "send",
			Method:      "POST",
			CurlExample: "curl -X POST -d @examples/send_email.json  http://0.0.0.0:1337/v1/gmail/send",
			Validator:   libemail.GenerateJSONValidator(PostSendGmailRequest{}),
		},
	}
}
func (e Gmail) HandlerById(i int) libemail.Service {
	switch i {
	// Add handlers for routes here
	case 0:
		return PostSendEmailsHandler
	}
	// This makes the server return a 404 by default
	return nil
}

func (e Gmail) LongPath(route libemail.Route) string {
	return libemail.DefaultLongPath(e, route)
}
