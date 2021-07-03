package core

import (
	"fmt"
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
	Call(method, namespace, apiName, queryStrings string)
	CallWithQueryStringGeneratorFunc(method, namespace, apiName string, queryStringGenerator func(p ...interface{}) string, params ...interface{})
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
)

func (c *core) Call(method, namespace, apiName, queryStrings string) {
	urlPath := c.createUrlPath(namespace, apiName)

	queryString := fmt.Sprintf(
		"%s&%s",
		c.createRequiredQueryString(),
		queryStrings,
	)

	fmt.Println(urlPath, queryString)
}

func (c *core) CallWithQueryStringGeneratorFunc(method, namespace, apiName string, queryStringGenerator func(p ...interface{}) string, params ...interface{}) {
	c.Call(method, namespace, apiName, queryStringGenerator(params...))
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
