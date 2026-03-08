package courses

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"

	"github.com/eWloYW8/zju-courses-go-sdk/zjuam"
)

var metaRefreshRe = regexp.MustCompile(`meta http-equiv="refresh" content="0;URL=([^"]+)"`)

// LoginWithZJUAM authenticates the courses client using a ZJUAM CAS client.
// This follows the full OAuth/CAS redirect chain between courses.zju.edu.cn and zjuam.zju.edu.cn.
func (c *Client) LoginWithZJUAM(ctx context.Context, am *zjuam.Client) error {
	// Create a no-redirect client sharing the same cookie jar as the SDK client
	noRedirect := &http.Client{
		Jar: c.sdk.HTTPClient().Jar,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	currentURL := "https://courses.zju.edu.cn/user/index"

	// Phase 1: Follow redirects from courses.zju.edu.cn until we reach zjuam.zju.edu.cn
	for {
		u, err := url.Parse(currentURL)
		if err != nil {
			return fmt.Errorf("courses: invalid URL %s: %w", currentURL, err)
		}
		if u.Hostname() == "zjuam.zju.edu.cn" {
			break
		}

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, currentURL, nil)
		if err != nil {
			return fmt.Errorf("courses: failed to create request: %w", err)
		}
		req.Header.Set("User-Agent", zjuam.UserAgent)

		resp, err := noRedirect.Do(req)
		if err != nil {
			return fmt.Errorf("courses: redirect failed at %s: %w", currentURL, err)
		}
		resp.Body.Close()

		location := resp.Header.Get("Location")
		if location == "" {
			return fmt.Errorf("courses: unexpected non-redirect response at %s (status %d)", currentURL, resp.StatusCode)
		}
		currentURL = location
	}

	// Phase 2: Authenticate via ZJUAM
	u, _ := url.Parse(currentURL)
	service := u.Query().Get("service")
	if service == "" {
		return fmt.Errorf("courses: no service parameter found in ZJUAM redirect URL")
	}

	ticketURL, err := am.LoginService(ctx, service)
	if err != nil {
		return fmt.Errorf("courses: ZJUAM authentication failed: %w", err)
	}

	// Phase 3: Follow the ticket URL back through courses.zju.edu.cn redirects
	currentURL = ticketURL
	for {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, currentURL, nil)
		if err != nil {
			return fmt.Errorf("courses: failed to create request: %w", err)
		}
		req.Header.Set("User-Agent", zjuam.UserAgent)

		resp, err := noRedirect.Do(req)
		if err != nil {
			return fmt.Errorf("courses: post-login redirect failed at %s: %w", currentURL, err)
		}

		// Handle meta http-equiv="refresh" redirects
		if resp.StatusCode == http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			resp.Body.Close()
			if matches := metaRefreshRe.FindSubmatch(body); len(matches) >= 2 {
				currentURL = string(matches[1])
				continue
			}
			break
		}

		resp.Body.Close()

		if resp.StatusCode >= 300 && resp.StatusCode < 400 {
			location := resp.Header.Get("Location")
			if location == "" {
				break
			}
			currentURL = location
			continue
		}

		break
	}

	return nil
}
