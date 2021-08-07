package model

import "github.com/jinzhu/gorm"

// MdCategory 表名和结构体名对应关系：大写开头驼峰转小写下划线驼峰
// 这里确认渲染没问题
//MdCategory is the category
type MdCategory struct {
	Id     int
	Pid    int
	Title  string `gorm:"size:30; unique"`
	Intro  string
	Icon   string
	Cnt    int
	Sort   int
	Status bool
}

// Get is to add the homepage
func (category *MdCategory) Get(db *gorm.DB, pid int, status bool) (cates []*MdCategory, err error) {
	if err := db.Order("-status, sort, title").Find(&cates).Error; err != nil {
		return nil, err
	}
	return cates, nil
}
