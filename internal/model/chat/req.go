package chat

type Req struct {
	Type int32       `json:"type"`
	Body interface{} `json:"body"`
}
