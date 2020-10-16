package libemail

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"io"
	"os"
)

func LoadToken(path string) (*oauth2.Token, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return TokenFromReader(file)
}

func TokenFromReader(r io.Reader) (*oauth2.Token, error) {
	token := &oauth2.Token{}
	err := json.NewDecoder(r).Decode(token)
	return token, err
}

// Saves a token to a file path.
func SaveToken(path string, token *oauth2.Token) error {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()
	return json.NewEncoder(f).Encode(token)
}

// Request a token from the web, then returns the retrieved token.
func GetTokenFromWeb(config *oauth2.Config) (*oauth2.Token, error) {
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", config.AuthCodeURL("state-token", oauth2.AccessTypeOffline))

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		return nil, fmt.Errorf("unable to read authorization code: %e", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		return nil, err
	}
	return tok, nil
}
