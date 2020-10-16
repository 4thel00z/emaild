package libemail

import (
	"encoding/base64"
	"encoding/json"
	"golang.org/x/oauth2"
)

type SmartString string
type Base64 SmartString

func (val *Base64) Unpack() (string, error) {
	decodeString, err := base64.StdEncoding.DecodeString(string(*val))
	return string(decodeString), err
}

func (val *Base64) UnmarshalJSON(b []byte) error {

	var raw string
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}

	decoded, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		return err
	}
	*val = Base64(decoded)
	return nil
}

// default *SmartString.String() to "" to avoid panicking
func (val *SmartString) String() string {
	if val == nil {
		return ""
	}
	return string(*val)
}

type Email struct {
	// Which of the accounts registered via the pool module you want to schedule on
	Account     string   `json:"account" validate:"empty=false"`
	To          []string `json:"to" validate:"empty=false"`
	From        string   `json:"from" validate:"empty=false"`
	Cc          []string `json:"cc"`
	Bcc         []string `json:"bcc"`
	Subject     string   `json:"subject"`
	ReplyTo     []string `json:"reply_to"`
	Sender      string   `json:"sender"`
	Attachments []string `json:"attachments"`
	//base64 encoded
	Body *string `json:"body,omitempty"`
	//base64 encoded
	HTML *string      `json:"html,omitempty"`
	File *SmartString `json:"file,omitempty"`
	// delay in seconds from now
	Delay int `json:"delay" validate:"gte=0"`
}

type Sender interface {
	Init(config *oauth2.Config, token *oauth2.Token) error
	Cleanup() error
	Send(message *Email) (interface{}, error)
}
