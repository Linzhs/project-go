package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type LogLevel uint8

const (
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelWarning
	LogLevelError
	LogLevelFatal
)

type Caller interface {
	Call(ctx context.Context) (interface{}, error)
}

type ClientOptFunc func(c *Client) error

type Client struct {
	err      error
	ctx      context.Context
	traceId  string
	logLevel LogLevel

	client *http.Client
	header http.Header

	url           string
	method        string
	callNameSpace string
	query         map[string]string
	body          []byte
}

func NewHttpClient(ctx context.Context, opts ...ClientOptFunc) (*Client, error) {
	if ctx == nil {
		return nil, fmt.Errorf("invalid context object ")
	}

	c := &Client{
		ctx:    ctx,
		client: http.DefaultClient,
		query:  map[string]string{},

		header: map[string][]string{},
	}

	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func SetTraceId(traceId string) ClientOptFunc {
	return func(c *Client) error {
		c.traceId = traceId
		return nil
	}
}

func SetTraceLogLevel(level LogLevel) ClientOptFunc {
	return func(c *Client) error {
		c.logLevel = level
		return nil
	}
}

func SetHttpClient(client *http.Client) ClientOptFunc {
	return func(c *Client) error {
		if c == nil {
			return errors.New("invalid http.Client, please check it! ")
		}

		c.client = client
		return nil
	}
}

func SetTimeout(timeout time.Duration) ClientOptFunc {
	return func(c *Client) error {
		c.client.Timeout = timeout
		return nil
	}
}

func SetUrl(url string) ClientOptFunc {
	return func(c *Client) error {
		c.url = url
		return nil
	}
}

func (c *Client) Ctx() context.Context {
	return c.ctx
}

func (c *Client) errNonNil() bool {
	return c.err != nil
}

func (c *Client) copyHeaderFromCtx() *Client {
	if c.errNonNil() {
		return c
	}

	//c.header.Add(rest.HeaderAcceptLanguage, utils.LangFromContext(c.Ctx()))
	//c.header.Add(rest.HeaderCurrency, utils.CurrencyFromContext(c.Ctx()))
	//c.header.Add(rest.HeaderToken, utils.TokenFromContext(c.Ctx()))
	//c.header.Add(rest.HeaderXKlookRequestID, utils.RequestIDFromContext(c.Ctx()))

	return c
}

func (c *Client) WithHeader(key, val string) *Client {
	if c.errNonNil() {
		return c
	}

	c.header.Add(key, val)

	return c
}

func (c *Client) WithUrl(url string) *Client {
	if c.errNonNil() {
		return c
	}

	c.url = url

	return c
}

func (c *Client) WithQuery(key, val string) *Client {
	if c.errNonNil() {
		return c
	}

	c.query[key] = val

	return c
}

func (c *Client) WithBody(b []byte) *Client {
	if c.errNonNil() {
		return c
	}

	c.body = b

	return c
}

func (c *Client) BodyWithJsonMarshal(v interface{}) *Client {
	if c.errNonNil() {
		return c
	}

	c.body, c.err = json.Marshal(v)

	return c
}

func (c *Client) WithRequestMethod(m string) *Client {
	if c.errNonNil() {
		return c
	}

	// 底层 http.NewRequest 默认填写为 GET
	c.method = m

	return c
}

func (c *Client) WithCallStatNameSpace(s string) *Client {
	if c.errNonNil() {
		return c
	}

	c.callNameSpace = s

	return c
}

func (c *Client) buildRequest() (*http.Request, error) {
	if c.errNonNil() {
		return nil, c.err
	}

	if len(c.url) == 0 {
		return nil, fmt.Errorf("invalid url format. ")
	}

	req, err := http.NewRequest(c.method, c.url, bytes.NewReader(c.body))
	if err != nil {
		return nil, err
	}

	if c.query != nil {
		q := req.URL.Query()
		for k, v := range c.query {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if c.header != nil {
		for key, s := range c.header {
			for _, val := range s {
				req.Header.Set(key, val)
			}
		}
	}

	return req, err
}

func (c *Client) Call() (response *http.Response, err error) {
	if c.errNonNil() {
		return nil, c.err
	}

	//bt := time.Now()
	defer func() {
		//collect(c.Ctx(), c.callNameSpace, err, time.Since(bt).Nanoseconds())
	}()

	req, err := c.buildRequest()
	if err != nil {
		return nil, fmt.Errorf("failed to build http request: %s", err)
	}

	if len(c.callNameSpace) == 0 {
		c.callNameSpace = req.URL.Path
	}

	return c.client.Do(req)
}

func (c *Client) PrintErr(level LogLevel) {
	if c.err == nil {
		return
	}

	switch level {
	case LogLevelDebug:
	case LogLevelInfo:
	case LogLevelWarning:
	case LogLevelError:
	case LogLevelFatal:
	default:

	}
}

func ErrNonNil(c *Client) bool {
	if c == nil {
		return false
	}

	return c.errNonNil()
}
