package database

import (
	"github.com/samarthasthan/twitter-sentiments/types"
	"gorm.io/gorm"
)

type Database interface {
	CreateTweet()
}

type MySqlDB struct {
	DB *gorm.DB
}

func NewMySqlDB(db *gorm.DB) *MySqlDB {
	return &MySqlDB{
		DB: db,
	}
}

func (db *MySqlDB) CreateTweet(tweets []*types.SentimentResult) error {
	result := db.DB.CreateInBatches(tweets, 10)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (db *MySqlDB) GetTweets(limit int, offset int) ([]types.SentimentResult, error) {
	var res []types.SentimentResult
	result := db.DB.Limit(limit).Offset(offset).Find(&res)
	if result.Error != nil {
		return res, result.Error
	}
	return res, nil
}
