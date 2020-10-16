package gmail

import (
	"emaild/pkg/libemail"
	"errors"
)

type PostSendGmailRequest libemail.Email

func (r PostSendGmailRequest) Validate() error {
	if r.Body == nil && r.HTML == nil {
		return errors.New("one of 'body' and 'html' fields should be non nil")
	}
	return nil
}
