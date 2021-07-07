package core

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type OutputFormat string

type core struct {
	resellerId   string
	apiKey       string
	isProduction bool
	outFormat    OutputFormat
}

type JSONStatusResponse struct {
	Status  string `json:"status"`
	Message string `jsob:"message"`
}

type Core interface {
	CallApi(method, namespace, apiName string, data url.Values) (*http.Response, error)
	PrintResponse(data []byte) error
}

const (
	OUTPUT_JSON OutputFormat = "json"
	OUTPUT_XML  OutputFormat = "xml"
)

var (
	host = map[bool]string{
		true:  "https://httpapi.com/api",
		false: "https://test.httpapi.com/api",
	}

	ErrRcApiUnsupportedMethod = errors.New("unsupported http method")
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
		"%s/%s/%s.%s",
		host[c.isProduction],
		namespace,
		apiName,
		string(c.outFormat),
	)
}

func (c *core) createRequiredQueryString() string {
	return fmt.Sprintf(
		"auth-userid=%s&api-key=%s",
		url.QueryEscape(c.resellerId),
		url.QueryEscape(c.apiKey),
	)
}

func New(resellerId, apiKey string, outFormat OutputFormat, isProduction bool) Core {
	return &core{
		resellerId:   resellerId,
		apiKey:       apiKey,
		isProduction: isProduction,
		outFormat:    outFormat,
	}
}
