package zjuam

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
)

const (
	pubkeyURL   = "https://zjuam.zju.edu.cn/cas/v2/getPubKey"
	casLoginURL = "https://zjuam.zju.edu.cn/cas/login"
	UserAgent   = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0"
)

var (
	executionRe = regexp.MustCompile(`name="execution" value="([^"]+)"`)
	msgRe       = regexp.MustCompile(`<span id="msg">([^<]+)</span>`)
)

// Client handles authentication with ZJU's CAS (zjuam.zju.edu.cn).
type Client struct {
	username   string
	password   string
	httpClient *http.Client
	loggedIn   bool
}

// NewClient creates a new ZJUAM authentication client.
func NewClient(username, password string) *Client {
	jar, _ := cookiejar.New(nil)
	return &Client{
		username: username,
		password: password,
		httpClient: &http.Client{
			Jar: jar,
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
	}
}

type pubKeyResponse struct {
	Modulus  string `json:"modulus"`
	Exponent string `json:"exponent"`
}

func (c *Client) doRequest(ctx context.Context, method, reqURL string, body io.Reader, contentType string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, reqURL, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", UserAgent)
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	return c.httpClient.Do(req)
}

func (c *Client) doGet(ctx context.Context, reqURL string) (*http.Response, error) {
	return c.doRequest(ctx, http.MethodGet, reqURL, nil, "")
}

func (c *Client) login(ctx context.Context, loginURL string) (string, error) {
	// Step 1: Get login page to extract execution token
	resp, err := c.doGet(ctx, loginURL)
	if err != nil {
		return "", fmt.Errorf("zjuam: failed to fetch login page: %w", err)
	}
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return "", fmt.Errorf("zjuam: failed to read login page: %w", err)
	}

	matches := executionRe.FindSubmatch(body)
	if len(matches) < 2 {
		return "", fmt.Errorf("zjuam: execution token not found in login page")
	}
	execution := string(matches[1])

	// Step 2: Fetch RSA public key
	resp, err = c.doGet(ctx, pubkeyURL)
	if err != nil {
		return "", fmt.Errorf("zjuam: failed to fetch public key: %w", err)
	}
	var pubkey pubKeyResponse
	err = json.NewDecoder(resp.Body).Decode(&pubkey)
	resp.Body.Close()
	if err != nil {
		return "", fmt.Errorf("zjuam: failed to decode public key: %w", err)
	}

	// Step 3: Encrypt password and POST login form
	encryptedPassword := rsaEncrypt(c.password, pubkey.Exponent, pubkey.Modulus)

	formData := url.Values{
		"username":  {c.username},
		"password":  {encryptedPassword},
		"execution": {execution},
		"_eventId":  {"submit"},
		"authcode":  {""},
	}

	resp, err = c.doRequest(ctx, http.MethodPost, loginURL,
		strings.NewReader(formData.Encode()), "application/x-www-form-urlencoded")
	if err != nil {
		return "", fmt.Errorf("zjuam: failed to post login: %w", err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusFound:
		c.loggedIn = true
		return resp.Header.Get("Location"), nil
	case http.StatusOK:
		body, _ := io.ReadAll(resp.Body)
		msg := "unknown error"
		if m := msgRe.FindSubmatch(body); len(m) >= 2 {
			msg = string(m[1])
		}
		return "", fmt.Errorf("zjuam: login failed: %s", msg)
	default:
		return "", fmt.Errorf("zjuam: login failed with status %d", resp.StatusCode)
	}
}

// Login performs CAS login to zjuam.zju.edu.cn.
func (c *Client) Login(ctx context.Context) error {
	_, err := c.login(ctx, casLoginURL)
	return err
}

// LoginService performs CAS login for a specific service URL and returns
// the redirect URL containing the service ticket.
func (c *Client) LoginService(ctx context.Context, service string) (string, error) {
	fullLoginURL := casLoginURL + "?service=" + url.QueryEscape(service)

	if c.loggedIn {
		resp, err := c.doGet(ctx, fullLoginURL)
		if err != nil {
			return "", fmt.Errorf("zjuam: failed to fetch service login: %w", err)
		}
		resp.Body.Close()

		switch resp.StatusCode {
		case http.StatusFound:
			return resp.Header.Get("Location"), nil
		case http.StatusOK:
			return c.login(ctx, fullLoginURL)
		default:
			return "", fmt.Errorf("zjuam: service login failed with status %d", resp.StatusCode)
		}
	}

	return c.login(ctx, fullLoginURL)
}

// Fetch makes an HTTP request using the ZJUAM session cookies.
// Automatically performs login if not already logged in.
func (c *Client) Fetch(ctx context.Context, reqURL string, opts ...RequestOption) (*http.Response, error) {
	if !c.loggedIn {
		if err := c.Login(ctx); err != nil {
			return nil, err
		}
	}

	o := &requestOptions{method: http.MethodGet}
	for _, opt := range opts {
		opt(o)
	}

	return c.doRequest(ctx, o.method, reqURL, o.body, o.contentType)
}

type requestOptions struct {
	method      string
	body        io.Reader
	contentType string
}

// RequestOption configures an HTTP request made via Fetch.
type RequestOption func(*requestOptions)

// WithMethod sets the HTTP method.
func WithMethod(method string) RequestOption {
	return func(o *requestOptions) {
		o.method = method
	}
}

// WithBody sets the request body and content type.
func WithBody(body io.Reader, contentType string) RequestOption {
	return func(o *requestOptions) {
		o.body = body
		o.contentType = contentType
	}
}
