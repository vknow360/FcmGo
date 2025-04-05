package types

type NotificationRequest struct {
	Title string            `json:"title"`
	Body  string            `json:"body"`
	Topic string            `json:"topic"`
	Sound string            `json:"sound"`
	Image string            `json:"image"`
	Data  map[string]string `json:"data"`
}

type TokenRequest struct {
	Title  string            `json:"title"`
	Body   string            `json:"body"`
	Tokens []string          `json:"tokens"`
	Sound  string            `json:"sound"`
	Image  string            `json:"image"`
	Data   map[string]string `json:"data"`
}
