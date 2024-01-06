package handler

import (
	"context"

	"github.com/samarthasthan/twitter-sentiments/database"
	pb "github.com/samarthasthan/twitter-sentiments/proto"
	"github.com/samarthasthan/twitter-sentiments/types"
)

type TweetGrpcServer struct {
	pb.UnimplementedTweetServiceServer
	DB *database.MySqlDB
}

func NewTweetGrpcServer(db *database.MySqlDB) *TweetGrpcServer {
	return &TweetGrpcServer{
		DB: db,
	}
}

func (s *TweetGrpcServer) TweetsHandler(ctx context.Context, in *pb.Pagination) (*pb.Tweets, error) {
	res, err := s.DB.GetTweets(int(in.GetLimit()), int(in.GetOffset()))
	if err != nil {
		return nil, err
	}
	pbTweets := convertToPbTweets(res)
	return pbTweets, nil
}

func convertToPbTweets(sentimentResults []types.SentimentResult) *pb.Tweets {
	var tweets []*pb.Tweet

	for _, result := range sentimentResults {
		var score int32
		if result.Score != nil {
			score = *result.Score
		}
		tweet := &pb.Tweet{
			Id:       int64(result.ID),
			Username: result.Username,
			Content:  result.Content,
			Score:    score,
		}
		tweets = append(tweets, tweet)
	}

	return &pb.Tweets{Tweets: tweets}
}
