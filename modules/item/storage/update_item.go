package storage

import (
	"Learn_Golang/modules/item/model"
	"context"
)

func (sql *sqlStore) UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *model.TodoItemUpdate) error {

	if err := sql.db.Where(cond).Updates(dataUpdate); err.Error != nil {
		return err.Error
	}
	return nil
}
