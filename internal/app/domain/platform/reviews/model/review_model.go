package model

import (
	"math"
	"time"
)

type Review struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	BookId    uint      `json:"bookId"`
	UserId    uint      `json:"userId"`
	Rating    int       `json:"rating" binding:"required"`
	Title     string    `json:"title"`
	Review    string    `json:"review"`
	AvgReview float64   `json:"avgReview"`
	CreatedAt time.Time `time_format:"2006-01-02 15:04:05"`
	UpdatedAt time.Time `time_format:"2006-01-02 15:04:05"`
}

func (r *Review) CalculateAvg(reviews []Review) {
	sum := 0
	for _, review := range reviews {
		sum += review.Rating
	}
	if sum == 0 {
		r.AvgReview = 0
	}
	avg := float64(sum) / float64(len(reviews))
	r.AvgReview = math.Round(avg*10) / 10
}
