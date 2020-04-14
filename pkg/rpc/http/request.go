package http

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type Request http.Request

func NewRequest(method, url string) (*Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	return (*Request)(req), nil
}

func (r *Request) SetBody(body interface{}, gzipCompress bool) error {
	switch b := body.(type) {
	case string:
		if gzipCompress {

		}
		return r.setBodyString(b)
	default:

	}

	return r.setBodyJson(body)
}

func (r *Request) setBodyJson(data interface{}) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}
	r.Header.Set("Content-Type", "application/json")
	_ = r.setBodyReader(bytes.NewReader(body))
	return nil
}

func (r *Request) setBodyString(body string) error {
	return r.setBodyReader(strings.NewReader(body))
}

func (r *Request) setBodyReader(body io.Reader) error {
	rc, ok := body.(io.ReadCloser)
	if !ok && body != nil {
		rc = ioutil.NopCloser(body)
	}
	r.Body = rc
	if body != nil {
		switch v := body.(type) {
		case *strings.Reader:
			r.ContentLength = int64(v.Len())
		case *bytes.Buffer:
			r.ContentLength = int64(v.Len())
		}
	}

	return nil
}
