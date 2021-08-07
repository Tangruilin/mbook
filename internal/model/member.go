package model

import "time"

type Member struct {
	MemberId      int       `gorm:"pk;auto" json:"member_id"`
	Account       string    `gorm:"size:30;unique" json:"account"`
	Nickname      string    `gorm:"size:30;unique" json:"nickname"`
	Password      string    ` json:"-"`
	Description   string    `gorm:"size:640" json:"description"`
	Email         string    `gorm:"size:100;unique" json:"email"`
	Phone         string    `gorm:"size:20:;null;default(null)" json:"phone"`
	Avatar        string    `json:"avatar"`
	Role          int       `gorm:"default:1" json:"role"`
	RoleName      string    `gorm:"-" json:"role_name"`
	Status        int       `gorm:"default:0" json:"status"`
	CreateTime    time.Time `gorm:"type:datetime;auto_now_add" json:"create_time"`
	CreateAt      int       `json:"create_at"`
	LastLoginTime time.Time `gorm:"type:datetime;null" json:"last_login_time"`
}

func NewMember() *Member {
	return &Member{}
}
