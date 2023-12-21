package entity

import (
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	PhoneNumber string `gorm:"uniqueIndex" valid:"stringlength(10|10)~PhoneNumber length is not 10 digits."`
	Password    string
	Url         string `valid:"url~Url is invalid"`
}
