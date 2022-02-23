package chat

type Msg struct {
	Type    uint8  `json:"type"`
	RoomIdx int    `json:"roomIdx"`
	UserIdx int    `json:"userIdx"`
	Name    string `json:"name"`
	Body    string `json:"body"`
}

