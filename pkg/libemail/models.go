package libemail

import (
	"encoding/json"
	"github.com/monzo/typhon"
	"io/ioutil"
	"os"
	"strings"
)

type GenericResponse struct {
	Message interface{} `json:"message"`
	Error   *string     `json:"error,omitempty"`
}

func ParseConfig(path string) (config Config, err error) {
	file, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return Config{}, err
	}
	err = json.Unmarshal(content, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

type Config struct {
	//TODO: add more fields here if you want to make the app more configurable
}

type Service func(app App) typhon.Service
type Validator func(request typhon.Request) error

type Route struct {
	Path        string `json:"path"`
	Method      string `json:"method"`
	CurlExample string `json:"curl_example"`
	LongPath    string
	Validator   *Validator
}

type Module interface {
	Version() string
	Namespace() string
	Routes() []Route
	HandlerById(int) Service
	LongPath(route Route) string
}

func DefaultLongPath(module Module, route Route) string {
	return "/" + strings.Join([]string{module.Version(), module.Namespace(), route.Path}, "/")
}