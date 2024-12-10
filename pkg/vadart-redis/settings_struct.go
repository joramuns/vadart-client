package vadart_redis

type Settings struct {
	Receiver string `json:"receiver"`
	Command  string `json:"command"`
	Value    string `json:"value"`
}
