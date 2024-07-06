package storage

import (
	"Learn_Golang/modules/item/model"
	"context"
)

func (sql *sqlStore) DeleteItem(ctx context.Context, cond map[string]interface{}) error {
	deleteStatus := model.ItemStatusDeleted
	if err := sql.db.Table(model.TodoItem{}.TableName()).Where(cond).Updates(map[string]interface{}{
		"status": deleteStatus.String(),
	}); err.Error != nil {
		return err.Error
	}
	return nil
}
