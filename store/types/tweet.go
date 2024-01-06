package types

type Tweet struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

type SentimentResult struct {
	ID       uint   `gorm:"primarykey"`
	Username string `json:"username"`
	Content  string `json:"content"`
	Score    *int32 `json:"score"`
}
