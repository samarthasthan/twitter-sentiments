package types

import "gorm.io/gorm"

type Tweet struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

type SentimentResult struct {
	gorm.Model
	Username string `json:"username"`
	Content  string `json:"content"`
	Score    int    `json:"score"`
}
