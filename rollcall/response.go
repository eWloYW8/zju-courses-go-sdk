package rollcall

import "encoding/json"

type CourseRollcallsResponse struct {
	Rollcalls []json.RawMessage `json:"rollcalls"`
}
