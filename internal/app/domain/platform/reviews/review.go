package reviews

import (
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"math"
	"time"
)

type Review struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	BookId    uint      `json:"bookId"`
	UserId    uint      `json:"userId"`
	Rating    int       `json:"rating" binding:"required"`
	AvgReview float64   `json:"avgReview"`
	Title     string    `json:"title"`
	Review    string    `json:"review"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (r Review) TableName() string {
	return "reviews"
}

type Module interface {
	Get(bookId string) ([]Review, *errors.ApiError)
	Find(id string) (Review, *errors.ApiError)
	Create(review Review) (*Review, *errors.ApiError)
	Update(id string, upReview Review) *errors.ApiError
	Delete(id string) *errors.ApiError
}

type Repository interface {
	Get(bookId string) ([]Review, *errors.ApiError)
	Find(id string) (Review, *errors.ApiError)
	Create(review Review) (*Review, *errors.ApiError)
	Update(id string, upReview Review) *errors.ApiError
	Delete(id string) *errors.ApiError
}

func (r *Review) CalculateAvg(reviews []Review) {
	sum := 0
	for i := 0; i < len(reviews); i++ {
		sum += reviews[i].Rating
	}
	if sum == 0 {
		r.AvgReview = 0
	}
	avg := float64(sum) / float64(len(reviews))
	r.AvgReview = math.Round(avg*10) / 10
}
