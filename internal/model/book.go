package model

import (
	"time"
)

type MdBooks struct {
	BookId         int       `gorm:"primary_key;" json:"book_id"`
	BookName       string    `gorm:"size:500" json:"book_name"`
	Identify       string    `gorm:"size:100" json:"identify"`
	OrderIndex     int       `gorm:"default:0" json:"order_index"`
	Description    string    `gorm:"size:1000" json:"description"`
	Cover          string    `gorm:"size:1000" json:"cover"`
	Editor         string    `gorm:"size:50" json:"editor"`
	Status         int       `gorm:"default:0" json:"status"`
	PrivatelyOwned int       `gorm:"default:0" json:"privately_owned"`
	PrivateToken   string    `gorm:"size:500" json:"private_token"`
	MemberId       int       `gorm:"size:100" json:"member_id"`
	CreatTime      time.Time `gorm:"type:datetime;auto_now_add" json:"creat_time"`
	ModifyTime     time.Time `gorm:"type:datetime;auto_now_add" json:"modify_time"`
	ReleaseTime    time.Time `gorm:"type:datetime;" json:"release_time"`
	DocCount       int       `json:"doc_count"`
	CommentCount   int       `gorm:"type:int" json:"comment_count"`
	Vcnt           int       `gorm:"default:0" json:"vcnt"`
	Collection     int       `gorm:"column:star;default:0" json:"collection"`
	Score          int       `gorm:"default:40" json:"score"`
	CntScore       int
	CntComment     int
	Author         string `gorm:"size:50"`
	AuthorUrl      string `gorm:"size:1000"`
}
