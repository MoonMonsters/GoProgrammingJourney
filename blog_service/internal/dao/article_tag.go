package dao

import "GoProgrammingJourney/blog_service/internal/model"

func (d *Dao) GetArticleTagByAID(articleId uint32) (model.ArticleTag, error) {
	articleTag := model.ArticleTag{
		ArticleId: articleId,
	}
	return articleTag.GetByAID(d.engine)
}

func (d *Dao) GetArticleTagListByTID(tagId uint32) ([]*model.ArticleTag, error) {
	articleTag := model.ArticleTag{
		TagId: tagId,
	}
	return articleTag.ListByTID(d.engine)
}

func (d *Dao) GetArticleTagListByAIDs(articleIDs []uint32) ([]*model.ArticleTag, error) {
	articleTag := model.ArticleTag{}
	return articleTag.ListByAIDs(d.engine, articleIDs)
}

func (d *Dao) CreateArticleTag(articleID, tagID uint32, createdBy string) error {
	articleTag := model.ArticleTag{
		Model: &model.Model{
			CreatedBy: createdBy,
		},
		ArticleId: articleID,
		TagId:     tagID,
	}
	return articleTag.Create(d.engine)
}

func (d *Dao) UpdateArticleTag(articleID, tagID uint32, modifiedBy string) error {
	articleTag := model.ArticleTag{
		ArticleId: articleID,
	}
	values := map[string]interface{}{
		"article_id":  articleID,
		"tag_id":      tagID,
		"modified_by": modifiedBy,
	}
	return articleTag.UpdateOne(d.engine, values)
}

func (d *Dao) DeleteArticleTag(articleID uint32) error {
	articleTag := model.ArticleTag{
		ArticleId: articleID,
	}
	return articleTag.DeleteOne(d.engine)
}
