package zjucourses

import (
	"net/http"

	"github.com/eWloYW8/zju-courses-go-sdk/activities"
	"github.com/eWloYW8/zju-courses-go-sdk/admin"
	"github.com/eWloYW8/zju-courses-go-sdk/aircredit"
	"github.com/eWloYW8/zju-courses-go-sdk/calendar"
	"github.com/eWloYW8/zju-courses-go-sdk/courses"
	"github.com/eWloYW8/zju-courses-go-sdk/exams"
	"github.com/eWloYW8/zju-courses-go-sdk/forum"
	"github.com/eWloYW8/zju-courses-go-sdk/homework"
	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
	"github.com/eWloYW8/zju-courses-go-sdk/notifications"
	"github.com/eWloYW8/zju-courses-go-sdk/resources"
	"github.com/eWloYW8/zju-courses-go-sdk/statistics"
	"github.com/eWloYW8/zju-courses-go-sdk/uploads"
	"github.com/eWloYW8/zju-courses-go-sdk/users"
)

const DefaultBaseURL = sdk.DefaultBaseURL

type Client struct {
	sdk *sdk.Client

	Courses       *courses.Service
	Activities    *activities.Service
	Homework      *homework.Service
	Forum         *forum.Service
	Exams         *exams.Service
	Uploads       *uploads.Service
	User          *users.Service
	Notifications *notifications.Service
	Calendar      *calendar.Service
	Statistics    *statistics.Service
	Admin         *admin.Service
	AirCredit     *aircredit.Service
	Resources     *resources.Service
}

type ClientOption = sdk.ClientOption

type Response = sdk.Response

type APIError = sdk.APIError

func WithHTTPClient(httpClient *http.Client) ClientOption {
	return sdk.WithHTTPClient(httpClient)
}

func WithBaseURL(baseURL string) ClientOption {
	return sdk.WithBaseURL(baseURL)
}

func WithCookies(cookies []*http.Cookie) ClientOption {
	return sdk.WithCookies(cookies)
}

func NewClient(opts ...ClientOption) *Client {
	core := sdk.NewClient(opts...)
	c := &Client{sdk: core}

	c.Courses = courses.New(core)
	c.Activities = activities.New(core)
	c.Homework = homework.New(core)
	c.Forum = forum.New(core)
	c.Exams = exams.New(core)
	c.Uploads = uploads.New(core)
	c.User = users.New(core)
	c.Notifications = notifications.New(core)
	c.Calendar = calendar.New(core)
	c.Statistics = statistics.New(core)
	c.Admin = admin.New(core)
	c.AirCredit = aircredit.New(core)
	c.Resources = resources.New(core)

	return c
}

func (c *Client) SetCookies(cookies []*http.Cookie) {
	c.sdk.SetCookies(cookies)
}

func (c *Client) SetCookieString(cookieStr string) {
	c.sdk.SetCookieString(cookieStr)
}
