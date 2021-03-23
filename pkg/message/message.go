package message

//Model is a Message Domain Model
type Model struct {
	ID       string `json:"id"`
	From     string `json:"from"`
	To       string `json:"to"`
	Msg      string `json:"msg"`
	DateTime string `json:"datetime"`
}
