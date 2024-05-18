package ports

type Email struct {
	To      string `json:"to"`
	From    string `json:"from"`
	Text    string `json:"text"`
	Subject string `json:"subject"`
}
