package biz

import (
	"Learn_Golang/modules/item/model"
	"context"
	"strings"
)

type CreateItemStore interface {
	CreateItem(ctx context.Context, data *model.TodoItemCreation) error
}
type crateItemBiz struct {
	store CreateItemStore
}

func NewCreateItemBiz(store CreateItemStore) *crateItemBiz {
	return &crateItemBiz{store: store}

}

func (biz *crateItemBiz) CreateItem(ctx context.Context, data *model.TodoItemCreation) error {
	title := strings.TrimSpace(data.Title)
	if title == "" {
		return model.ErrTiltleIsBlank
	}
	if err := biz.store.CreateItem(ctx, data); err != nil {
		return err
	}
	return nil

}
