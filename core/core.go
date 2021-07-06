package core

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type OutputFormat string

type core struct {
	resellerId   string
	apiKey       string
	isProduction bool
	outFormat    OutputFormat
}

type Core interface {
	CallApi(method, namespace, apiName string, data url.Values) (*http.Response, error)
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

func (c *core) CallApi(method, namespace, apiName string, data url.Values) (*http.Response, error) {
	urlPath := c.createUrlPath(namespace, apiName)
	switch method {
	case http.MethodGet:
		return http.Get(fmt.Sprintf("%s?%s&%s", urlPath, c.createRequiredQueryString(), data.Encode()))
	case http.MethodPost:
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
