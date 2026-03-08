package sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

const DefaultBaseURL = "https://courses.zju.edu.cn"

type Client struct {
	httpClient *http.Client
	baseURL    *url.URL
}

type ClientOption func(*Client)

func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) {
		u, err := url.Parse(baseURL)
		if err == nil {
			c.baseURL = u
		}
	}
}

func WithCookies(cookies []*http.Cookie) ClientOption {
	return func(c *Client) {
		if c.httpClient.Jar == nil {
			jar, _ := cookiejar.New(nil)
			c.httpClient.Jar = jar
		}
		c.httpClient.Jar.SetCookies(c.baseURL, cookies)
	}
}

func NewClient(opts ...ClientOption) *Client {
	jar, _ := cookiejar.New(nil)
	baseURL, _ := url.Parse(DefaultBaseURL)

	c := &Client{
		httpClient: &http.Client{Jar: jar},
		baseURL:    baseURL,
	}

	for _, opt := range opts {
		opt(c)
	}

	if c.httpClient.Jar == nil {
		c.httpClient.Jar = jar
	}

	return c
}

func (c *Client) SetCookies(cookies []*http.Cookie) {
	c.httpClient.Jar.SetCookies(c.baseURL, cookies)
}

func (c *Client) SetCookieString(cookieStr string) {
	header := http.Header{}
	header.Add("Cookie", cookieStr)
	request := http.Request{Header: header}
	cookies := request.Cookies()
	c.httpClient.Jar.SetCookies(c.baseURL, cookies)
}

func (c *Client) BaseURL() *url.URL {
	return c.baseURL
}

func (c *Client) NewRequest(ctx context.Context, method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasPrefix(urlStr, "http") {
		u, err := c.baseURL.Parse(urlStr)
		if err != nil {
			return nil, err
		}
		urlStr = u.String()
	}

	var buf io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, urlStr, buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")

	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response := &Response{Response: resp}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return response, &APIError{
			StatusCode: resp.StatusCode,
			Message:    string(bodyBytes),
			URL:        req.URL.String(),
		}
	}

	if v != nil {
		if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
			if err == io.EOF {
				return response, nil
			}
			return response, err
		}
	}

	return response, nil
}

func (c *Client) DoBytes(req *http.Request) (*Response, []byte, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	response := &Response{Response: resp}
	bodyBytes, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return response, nil, readErr
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return response, bodyBytes, &APIError{
			StatusCode: resp.StatusCode,
			Message:    string(bodyBytes),
			URL:        req.URL.String(),
		}
	}

	return response, bodyBytes, nil
}

func (c *Client) Get(ctx context.Context, url string, v interface{}) (*Response, error) {
	req, err := c.NewRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req, v)
}

func (c *Client) Post(ctx context.Context, url string, body, v interface{}) (*Response, error) {
	req, err := c.NewRequest(ctx, http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	return c.Do(req, v)
}

func (c *Client) Put(ctx context.Context, url string, body, v interface{}) (*Response, error) {
	req, err := c.NewRequest(ctx, http.MethodPut, url, body)
	if err != nil {
		return nil, err
	}
	return c.Do(req, v)
}

func (c *Client) Patch(ctx context.Context, url string, body, v interface{}) (*Response, error) {
	req, err := c.NewRequest(ctx, http.MethodPatch, url, body)
	if err != nil {
		return nil, err
	}
	return c.Do(req, v)
}

func (c *Client) Delete(ctx context.Context, url string, v interface{}) (*Response, error) {
	req, err := c.NewRequest(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req, v)
}

func (c *Client) DeleteWithBody(ctx context.Context, url string, body, v interface{}) (*Response, error) {
	req, err := c.NewRequest(ctx, http.MethodDelete, url, body)
	if err != nil {
		return nil, err
	}
	return c.Do(req, v)
}

func AddListOptions(urlStr string, page, pageSize int) string {
	sep := "?"
	if strings.Contains(urlStr, "?") {
		sep = "&"
	}
	if page > 0 {
		urlStr += fmt.Sprintf("%spage=%d", sep, page)
		sep = "&"
	}
	if pageSize > 0 {
		urlStr += fmt.Sprintf("%spage_size=%d", sep, pageSize)
	}
	return urlStr
}

func AddQueryParams(urlStr string, params map[string]string) string {
	if len(params) == 0 {
		return urlStr
	}
	sep := "?"
	if strings.Contains(urlStr, "?") {
		sep = "&"
	}
	for k, v := range params {
		urlStr += fmt.Sprintf("%s%s=%s", sep, url.QueryEscape(k), url.QueryEscape(v))
		sep = "&"
	}
	return urlStr
}
