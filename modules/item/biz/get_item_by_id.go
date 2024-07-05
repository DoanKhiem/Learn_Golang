package biz

import (
	"Learn_Golang/modules/item/model"
	"context"
)

type GetItemStore interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
}
type getItemBiz struct {
	store GetItemStore
}

func NewGetItemBiz(store GetItemStore) *getItemBiz {
	return &getItemBiz{store: store}

}

func (biz *getItemBiz) GetItemById(ctx context.Context, id int) (*model.TodoItem, error) {
	data, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}
	return data, nil

}
