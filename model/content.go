package model

import "github.com/jinzhu/gorm"

type Content struct {
	gorm.Model
	Hash      string  `gorm:"UNIQUE;NOT NULL"`
	Account   Account `gorm:"FOREIGNKEY:id;ASSOCIATION_FOREIGNKEY:account_id;"`
	FromSite  string
	AccountId uint
	Content   string
	Review    int
}

var ContentModel = Content{}

/*
链表查询content表的中的内容，和作者
 */

func (ContentModel *Content) ContentIndex()(err string,content []Content)  {

	db.Preload("Account").Order("id desc").Limit(50).Find(&content)
	if  len(content) == 0 {
		return "没有数据",nil
	}
	return "",content
}
