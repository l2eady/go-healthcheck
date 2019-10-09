package lhttp

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type HttpCaller interface {
	GET() (*http.Response, error)
	POST() (*http.Response, error)
	PUT() (*http.Response, error)
	PATCH() (*http.Response, error)
	DELETE() (*http.Response, error)
	SetClient(client *http.Client)
	SetURL(url string)
	SetBody(body interface{})
	SetHeader(header map[string]string)
}

type Caller struct {
	URL      string
	Body     interface{}
	Header   map[string]string
	Client   *http.Client
	Response interface{}
}
type value struct {
	URL      string
	Method   string
	Body     interface{}
	Header   map[string]string
	Client   *http.Client
	Response interface{}
}

// New create the new http caller with parameters request url, request body, request http header, and response format
func New(url string, body interface{}, header map[string]string, resp interface{}) HttpCaller {
	return &Caller{
		URL:      url,
		Body:     body,
		Header:   header,
		Response: resp,
	}
}

// SetClient set new http client
func (c *Caller) SetClient(client *http.Client) {
	c.Client = client
}

// GET func call server with method GET
func (c *Caller) GET() (*http.Response, error) {
	return invoke(value{
		URL:      c.URL,
		Method:   http.MethodGet,
		Body:     c.Body,
		Header:   c.Header,
		Client:   c.Client,
		Response: c.Response,
	})
}

// POST func call server with method POST
func (c *Caller) POST() (*http.Response, error) {
	return invoke(value{
		URL:      c.URL,
		Method:   http.MethodPost,
		Body:     c.Body,
		Header:   c.Header,
		Client:   c.Client,
		Response: c.Response,
	})
}

// PUT func call server with method PUT
func (c *Caller) PUT() (*http.Response, error) {
	return invoke(value{
		URL:      c.URL,
		Method:   http.MethodPut,
		Body:     c.Body,
		Header:   c.Header,
		Client:   c.Client,
		Response: c.Response,
	})
}

// PATCH func call server with method PATCH
func (c *Caller) PATCH() (*http.Response, error) {
	return invoke(value{
		URL:      c.URL,
		Method:   http.MethodPatch,
		Body:     c.Body,
		Header:   c.Header,
		Client:   c.Client,
		Response: c.Response,
	})
}

// DELETE func call server with method DELETE
func (c *Caller) DELETE() (*http.Response, error) {
	return invoke(value{
		URL:      c.URL,
		Method:   http.MethodDelete,
		Body:     c.Body,
		Header:   c.Header,
		Client:   c.Client,
		Response: c.Response,
	})
}

func (c *Caller) SetURL(url string) {
	c.URL = url
}
func (c *Caller) SetBody(body interface{}) {
	c.Body = body
}
func (c *Caller) SetHeader(header map[string]string) {
	c.Header = header
}

func invoke(v value) (*http.Response, error) {

	if v.Client == nil {
		v.Client = &http.Client{}
	}

	buf := bytes.NewBuffer(nil)
	if v.Body != nil {
		err := json.NewEncoder(buf).Encode(v.Body)
		if err != nil {
			return nil, err
		}
	}
	// make a new http request
	req, err := http.NewRequest(v.Method, v.URL, buf)
	if err != nil {
		return nil, err
	}
	// set the request header
	for k, v := range v.Header {
		req.Header.Set(k, v)
	}

	resp, err := v.Client.Do(req)
	// cannot connect to server (request timeout)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(v.Response)

	return resp, err
}
