package libemail

import (
	"bytes"
	"encoding/json"
	"github.com/monzo/typhon"
	"gopkg.in/dealancer/validate.v2"
	"io/ioutil"
	"reflect"
)

func Default404Handler(app App) typhon.Service {
	return func(req typhon.Request) typhon.Response {
		// TODO: Change this body to a default 404 page
		response := req.Response(nil)
		response.StatusCode = 404
		return response
	}
}

func GenerateJSONValidator(i interface{}) *Validator {
	t := reflect.TypeOf(i)
	toValidate := reflect.New(t).Interface()

	validator := func(r typhon.Request) error {

		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}

		// As if nothing has ever happened .. ( ͡° ͜ʖ ͡°)
		r.Body = ioutil.NopCloser(bytes.NewReader(content))

		err = json.Unmarshal(content, &toValidate)

		if err != nil {
			return err
		}

		err = validate.Validate(toValidate)

		if err != nil {
			return err
		}

		return nil
	}

	return (*Validator)(&validator)
}
