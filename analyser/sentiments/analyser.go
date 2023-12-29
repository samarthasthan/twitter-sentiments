package sentiments

import (
	"log"

	"github.com/cdipaolo/sentiment"
)

type Analyser struct {
	model sentiment.Models
}

func NewAnalyser() *Analyser {
	model, err := sentiment.Restore()
	if err != nil {
		log.Fatalln(err)
	}
	return &Analyser{
		model: model,
	}
}

func (a *Analyser) Analyse(content string) int {
	analysis := a.model.SentimentAnalysis(content, sentiment.English)
	return int(analysis.Score)
}
