package model

import "github.com/jinzhu/gorm"

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

func (a Article) TableName() string {
	return "blog_article"
}

func (a Article) Create(db *gorm.DB) (*Article, error) {
	if err := db.Create(&a).Error; err != nil {
		return nil, err
	}
	return &a, nil
}

func (a Article) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(&a).Updates(values).Where("id = ? AND is_del = ?", a.ID).Error; err != nil {
		return err
	}

	return nil
}

func (a Article) Get(db *gorm.DB) (Article, error) {
	var article Article
	// 查询语句
	db = db.Where("id = ? AND state = ? AND is_del = ?", a.ID, a.State, 0)
	// 获取第一条数据, 并放入article实例中
	err := db.First(&article).Error
	// 数据未找到
	if err != nil && err != gorm.ErrRecordNotFound {
		return article, err
	}
	return article, nil
}

func (a Article) Delete(db *gorm.DB) error {
	if err := db.Where("id = ? AND is_del = ?", a.Model.ID, 0).Delete(&a).Error; err != nil {
		return err
	}
	return nil
}
