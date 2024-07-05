package storage

import (
	"Learn_Golang/modules/item/model"
	"context"
)

func (sql *sqlStore) CreateItem(ctx context.Context, data *model.TodoItemCreation) error {
	if err := sql.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
