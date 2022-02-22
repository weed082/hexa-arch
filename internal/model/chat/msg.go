package chat

type Msg struct {
	Type    uint8  `json:"type"`
	RoomIdx int    `json:"roomIdx"`
	UserIdx int    `json:"userIdx"`
	Name    string `json:"name"`
	Body    string `json:"body"`
}

func (m *Msg) GetRoomIdx() int {
	return m.RoomIdx
}

func (m *Msg) GetUserIdx() int {
	return m.UserIdx
}
