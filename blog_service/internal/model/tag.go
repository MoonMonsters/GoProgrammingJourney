package model

import "github.com/jinzhu/gorm"

type Tag struct {
	*Model
	// 标签名称
	Name string `json:"name"`
	// 状态
	State uint8 `json:"state"`
}

// 获取某类型标签数量
func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int

	// 过滤条件
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// 获取标签列表
func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

// 删除标签
func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

// 更新标签数据
// 更新所选字段
func (t Tag) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&t).Where("id = ? AND is_del = ?", t.ID, 0).Updates(values).Error
}

// 根据id删除数据
func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", t.Model.ID, 0).Delete(&t).Error
}
