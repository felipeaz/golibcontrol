package reviews

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"math"
	"time"
)

type Review struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	BookId    uint      `json:"bookId"`
	UserId    uint      `json:"userId"`
	Rating    int       `json:"rating" binding:"required"`
	AvgReview float64   `json:"avgReview"`
	Title     string    `json:"title"`
	Review    string    `json:"review"`
	CreatedAt time.Time `time_format:"2006-01-02 15:04:05"`
	UpdatedAt time.Time `time_format:"2006-01-02 15:04:05"`
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
