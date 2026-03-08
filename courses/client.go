package courses

import (
	"net/http"

	"github.com/eWloYW8/zju-courses-go-sdk/courses/activities"
	"github.com/eWloYW8/zju-courses-go-sdk/courses/admin"
	"github.com/eWloYW8/zju-courses-go-sdk/courses/aircredit"
	"github.com/eWloYW8/zju-courses-go-sdk/courses/calendar"
	"github.com/eWloYW8/zju-courses-go-sdk/courses/courses"
	"github.com/eWloYW8/zju-courses-go-sdk/courses/exams"
	"github.com/eWloYW8/zju-courses-go-sdk/courses/feedback"
	"github.com/eWloYW8/zju-courses-go-sdk/courses/forum"
	"github.com/eWloYW8/zju-courses-go-sdk/courses/groups"
	"github.com/eWloYW8/zju-courses-go-sdk/courses/homework"
	"github.com/eWloYW8/zju-courses-go-sdk/courses/interactions"
	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
	"github.com/eWloYW8/zju-courses-go-sdk/courses/knowledge"
	"github.com/eWloYW8/zju-courses-go-sdk/courses/meetings"
	"github.com/eWloYW8/zju-courses-go-sdk/courses/notifications"
	"github.com/eWloYW8/zju-courses-go-sdk/courses/others"
	"github.com/eWloYW8/zju-courses-go-sdk/courses/resources"
	"github.com/eWloYW8/zju-courses-go-sdk/courses/rollcall"
	"github.com/eWloYW8/zju-courses-go-sdk/courses/statistics"
	"github.com/eWloYW8/zju-courses-go-sdk/courses/syllabus"
	"github.com/eWloYW8/zju-courses-go-sdk/courses/uploads"
	"github.com/eWloYW8/zju-courses-go-sdk/courses/users"
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
	Users         *users.Service
	Notifications *notifications.Service
	Calendar      *calendar.Service
	Statistics    *statistics.Service
	Admin         *admin.Service
	AirCredit     *aircredit.Service
	Resources     *resources.Service
	Knowledge     *knowledge.Service
	Meetings      *meetings.Service
	Groups        *groups.Service
	Rollcall      *rollcall.Service
	Interactions  *interactions.Service
	Feedback      *feedback.Service
	Syllabus      *syllabus.Service
	Others        *others.Service
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
	c.Users = users.New(core)
	c.Notifications = notifications.New(core)
	c.Calendar = calendar.New(core)
	c.Statistics = statistics.New(core)
	c.Admin = admin.New(core)
	c.AirCredit = aircredit.New(core)
	c.Resources = resources.New(core)
	c.Knowledge = knowledge.New(core)
	c.Meetings = meetings.New(core)
	c.Groups = groups.New(core)
	c.Rollcall = rollcall.New(core)
	c.Interactions = interactions.New(core)
	c.Feedback = feedback.New(core)
	c.Syllabus = syllabus.New(core)
	c.Others = others.New(core)

	return c
}

func (c *Client) SetCookies(cookies []*http.Cookie) {
	c.sdk.SetCookies(cookies)
}

func (c *Client) SetCookieString(cookieStr string) {
	c.sdk.SetCookieString(cookieStr)
}
