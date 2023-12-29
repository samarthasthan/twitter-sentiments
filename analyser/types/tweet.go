package types

type Tweet struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

type SentimentResult struct {
	Username string `json:"username"`
	Content  string `json:"content"`
	Score    int    `json:"score"`
}
