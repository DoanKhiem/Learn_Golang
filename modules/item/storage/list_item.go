package storage

import (
	"Learn_Golang/common"
	"Learn_Golang/modules/item/model"
	"context"
)

func (sql *sqlStore) ListItem(ctx context.Context, filter *model.Filter, paging *common.Paging, moreKeys ...string) ([]model.TodoItem, error) {
	var result []model.TodoItem

	db := sql.db.Where("status <> ?", "Deleted").Where(filter)

	if f := filter; f != nil {
		if f.Status != "" {
			db = db.Where("status = ?", f.Status)
		}
	}

	if err := db.Table(model.TodoItem{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.Order("id desc").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
