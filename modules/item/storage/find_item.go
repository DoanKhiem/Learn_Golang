package storage

import (
	"Learn_Golang/modules/item/model"
	"context"
)

func (sql *sqlStore) GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error) {
	var data model.TodoItem
	if err := sql.db.Where(cond).First(&data); err.Error != nil {
		return nil, err.Error
	}
	return &data, nil
}
