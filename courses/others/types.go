package others

type LessonRoom struct {
	RoomCode string `json:"room_code,omitempty"`
	RoomName string `json:"room_name,omitempty"`
	AppID    any    `json:"app_id,omitempty"`
}

type RoomLocation struct {
	ID       int    `json:"id"`
	Building string `json:"building,omitempty"`
	RoomName string `json:"room_name,omitempty"`
	RoomCode string `json:"room_code,omitempty"`
	Seats    int    `json:"seats,omitempty"`
}
