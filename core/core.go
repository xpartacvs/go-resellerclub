package core

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"regexp"
)

type core struct {
	resellerId   string
	apiKey       string
	isProduction bool
}

type JSONStatusResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Core interface {
	CallApi(method, namespace, apiName string, data url.Values) (*http.Response, error)
	PrintResponse(data []byte) error
}

var (
	host = map[bool]string{
		true:  "https://httpapi.com/api",
		false: "https://test.httpapi.com/api",
	}

	RgxEmail  = regexp.MustCompile("^(?:(?:(?:(?:[a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(?:\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|(?:(?:\\x22)(?:(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(?:\\x20|\\x09)+)?(?:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(\\x20|\\x09)+)?(?:\\x22))))@(?:(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$")
	RgxNumber = regexp.MustCompile("^[0-9]+$")

	ErrRcApiUnsupportedMethod = errors.New("unsupported http method")
	ErrRcOperationFailed      = errors.New("operation failed")
	ErrRcInvalidCredential    = errors.New("invalid credential")
)

func (c *core) PrintResponse(data []byte) error {
	var buffer bytes.Buffer
	if err := json.Indent(&buffer, data, "", "\t"); err != nil {
		return err
	}
	buffer.WriteTo(os.Stdout)
	return nil
}

func (c *core) CallApi(method, namespace, apiName string, data url.Values) (*http.Response, error) {
	urlPath := c.createUrlPath(namespace, apiName)
	switch method {
	case http.MethodGet:
		return http.Get(fmt.Sprintf("%s?%s&%s", urlPath, c.createRequiredQueryString(), data.Encode()))
	case http.MethodPost:
		data.Add("auth-userid", c.resellerId)
		data.Add("api-key", c.apiKey)
		return http.PostForm(urlPath, data)
	}
	return nil, ErrRcApiUnsupportedMethod
}

func (c *core) createUrlPath(namespace, apiName string) string {
	return fmt.Sprintf(
		"%s/%s/%s.json",
		host[c.isProduction],
		namespace,
		apiName,
	)
}

func (c *core) createRequiredQueryString() string {
	return fmt.Sprintf(
		"auth-userid=%s&api-key=%s",
		url.QueryEscape(c.resellerId),
		url.QueryEscape(c.apiKey),
	)
}

func New(resellerId, apiKey string, isProduction bool) Core {
	return &core{
		resellerId:   resellerId,
		apiKey:       apiKey,
		isProduction: isProduction,
	}
}
