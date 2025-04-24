package models

import (
	"boilerplate/db"
	"boilerplate/forms"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	UserID  int64  `gorm:"column:user_id" json:"user_id"`
	Title   string `gorm:"column:title" json:"title"`
	Content string `gorm:"column:content; type:varchar(255); not null" json:"content"`
	User    User   `gorm:"foreignKey:UserID"`
}

func (Article) TableName() string {
	return "article"
}

type ArticleModel struct{}

func (m ArticleModel) Create(userID int64, form forms.CreateArticleForm) (articleID int64, err error) {

	article := Article{
		UserID:  userID,
		Title:   form.Title,
		Content: form.Content,
	}
	err = db.GetDB().Create(&article).Error
	return articleID, err
}

func (m ArticleModel) FindById(userID, id int64) (article Article, err error) {

	if err := db.GetDB().Model(&article).Where("user_id = ?", userID).Find(&article).Error; err != nil {
		return article, nil
	}
	return article, err
}

func (m ArticleModel) FindAll(userID int64) (articles []Article, err error) {
	if err := db.GetDB().Model(&articles).Where("user_id = ?", userID).Find(&articles).Error; err != nil {
		return []Article{}, nil
	}
	return articles, err
}

func (m ArticleModel) Update(userID int64, id int64, form forms.CreateArticleForm) (err error) {
	//METHOD 1
	//Check the article by ID using this way
	// _, err = m.One(userID, id)
	// if err != nil {
	// 	return err
	// }
	var article Article
	if err := db.GetDB().Where("user_id = ? AND id = ?", userID, id).Find(&article).Error; err != nil {
		return nil
	}
	article.Content = form.Content
	article.Title = form.Title
	if err = db.GetDB().Save(&article).Error; err != nil {
		return nil
	}

	return err
}

//Delete ...
func (m ArticleModel) Delete(userID, id int64) (err error) {

	if err := db.GetDB().Delete(&Article{}, id).Error; err != nil {
		return nil
	}

	return err
}
