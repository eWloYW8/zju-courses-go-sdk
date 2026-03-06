package zjucourses

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

	"github.com/eWloYW8/zju-courses-go-sdk/model"
)

const (
	DefaultBaseURL = "https://courses.zju.edu.cn"
)

// Client manages communication with the ZJU Courses API.
type Client struct {
	httpClient *http.Client
	baseURL    *url.URL

	Courses       *CoursesService
	Activities    *ActivitiesService
	Homework      *HomeworkService
	Forum         *ForumService
	Exams         *ExamsService
	Uploads       *UploadsService
	User          *UserService
	Notifications *NotificationsService
	Calendar      *CalendarService
	Statistics    *StatisticsService
	Admin         *AdminService
	AirCredit     *AirCreditService
	Resources     *ResourcesService
}

// ClientOption configures the Client.
type ClientOption func(*Client)

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// WithBaseURL sets a custom base URL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) {
		u, err := url.Parse(baseURL)
		if err == nil {
			c.baseURL = u
		}
	}
}

// WithCookies sets cookies for authentication.
// The ZJU Courses platform uses cookie-based authentication.
// Note: This option must be applied after the HTTP client has a cookie jar.
func WithCookies(cookies []*http.Cookie) ClientOption {
	return func(c *Client) {
		if c.httpClient.Jar == nil {
			jar, _ := cookiejar.New(nil)
			c.httpClient.Jar = jar
		}
		c.httpClient.Jar.SetCookies(c.baseURL, cookies)
	}
}

// NewClient creates a new ZJU Courses API client.
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

	// Ensure the HTTP client has a cookie jar
	if c.httpClient.Jar == nil {
		c.httpClient.Jar = jar
	}

	c.Courses = &CoursesService{client: c}
	c.Activities = &ActivitiesService{client: c}
	c.Homework = &HomeworkService{client: c}
	c.Forum = &ForumService{client: c}
	c.Exams = &ExamsService{client: c}
	c.Uploads = &UploadsService{client: c}
	c.User = &UserService{client: c}
	c.Notifications = &NotificationsService{client: c}
	c.Calendar = &CalendarService{client: c}
	c.Statistics = &StatisticsService{client: c}
	c.Admin = &AdminService{client: c}
	c.AirCredit = &AirCreditService{client: c}
	c.Resources = &ResourcesService{client: c}

	return c
}

// SetCookies sets authentication cookies on the client.
func (c *Client) SetCookies(cookies []*http.Cookie) {
	c.httpClient.Jar.SetCookies(c.baseURL, cookies)
}

// SetCookieString parses a cookie header string and sets cookies.
// Format: "name1=value1; name2=value2"
func (c *Client) SetCookieString(cookieStr string) {
	header := http.Header{}
	header.Add("Cookie", cookieStr)
	request := http.Request{Header: header}
	cookies := request.Cookies()
	c.httpClient.Jar.SetCookies(c.baseURL, cookies)
}

// Response wraps an HTTP response.
type Response struct {
	*http.Response
}

// newRequest creates an API request.
func (c *Client) newRequest(ctx context.Context, method, urlStr string, body interface{}) (*http.Request, error) {
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

// do sends an API request and decodes the JSON response into v.
func (c *Client) do(req *http.Request, v interface{}) (*Response, error) {
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

// get performs a GET request.
func (c *Client) get(ctx context.Context, url string, v interface{}) (*Response, error) {
	req, err := c.newRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return c.do(req, v)
}

// post performs a POST request.
func (c *Client) post(ctx context.Context, url string, body, v interface{}) (*Response, error) {
	req, err := c.newRequest(ctx, http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	return c.do(req, v)
}

// put performs a PUT request.
func (c *Client) put(ctx context.Context, url string, body, v interface{}) (*Response, error) {
	req, err := c.newRequest(ctx, http.MethodPut, url, body)
	if err != nil {
		return nil, err
	}
	return c.do(req, v)
}

// patch performs a PATCH request.
func (c *Client) patch(ctx context.Context, url string, body, v interface{}) (*Response, error) {
	req, err := c.newRequest(ctx, http.MethodPatch, url, body)
	if err != nil {
		return nil, err
	}
	return c.do(req, v)
}

// delete performs a DELETE request.
func (c *Client) delete(ctx context.Context, url string, v interface{}) (*Response, error) {
	req, err := c.newRequest(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}
	return c.do(req, v)
}

// APIError represents an error returned by the API.
type APIError struct {
	StatusCode int
	Message    string
	URL        string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API error %d: %s (URL: %s)", e.StatusCode, e.Message, e.URL)
}

// addListOptions adds pagination query params to a URL string.
func addListOptions(urlStr string, opts *model.ListOptions) string {
	if opts == nil {
		return urlStr
	}
	sep := "?"
	if strings.Contains(urlStr, "?") {
		sep = "&"
	}
	if opts.Page > 0 {
		urlStr += fmt.Sprintf("%spage=%d", sep, opts.Page)
		sep = "&"
	}
	if opts.PageSize > 0 {
		urlStr += fmt.Sprintf("%spage_size=%d", sep, opts.PageSize)
	}
	return urlStr
}

// addQueryParams appends query parameters from a map to a URL string.
func addQueryParams(urlStr string, params map[string]string) string {
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
