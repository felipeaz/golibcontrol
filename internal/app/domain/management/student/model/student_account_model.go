package model

type StudentAccount struct {
	Email          string `json:"email" binding:"required"`
	Password       string `json:"password" binding:"required"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Phone          string `json:"phone"`
	StudentAccount bool   `json:"studentAccount" gorm:"<-:create"`
}

type AccountResponse struct {
	ID uint `json:"id"`
}
