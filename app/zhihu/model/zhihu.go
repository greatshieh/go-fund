package model

import (
	"fmt"
	"gospider/global"
)

type ZhihuModel struct {
	Tag           *string `json:"tag" gorm:"column:tag" excel:"话题"`
	QuestionTitle *string `json:"title" gorm:"column:question_title" excel:"问题"`
	UpvoteCount   *int    `json:"voteup_count" gorm:"column:upvote_count" excel:"点赞数"`
	CommentCount  *int    `json:"comment_count" gorm:"column:comment_count" excel:"评论数"`
	AnswerLink    *string `json:"answer_link" gorm:"index;primaryKey;column:answer_link" excel:"链接"`
}

func (ZhihuModel) TableName() string {
	return "zhihu_topic"
}

func (t ZhihuModel) Save() {
	instance := new(ZhihuModel)
	if err := global.GPA_DB.Where("answer_link = ?", t.AnswerLink).First(&instance).Error; err == nil {
		// 找到记录并且更新记录
		global.GPA_DB.Model(instance).Where("answer_link = ?", t.AnswerLink).Updates(t)
		fmt.Printf("更新 ❗ %s, %s, %s\n", *t.Tag, *t.QuestionTitle, *t.AnswerLink)
	} else {
		// 没有找到记录
		global.GPA_DB.Create(&t)
		fmt.Printf("创建 ⚙ %s, %s, %s\n", *t.Tag, *t.QuestionTitle, *t.AnswerLink)
	}
}
