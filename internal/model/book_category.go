package model

type MdBookCategory struct {
	Id         int `gorm:"primary_key"`
	BookId     int
	CategoryId int
}
