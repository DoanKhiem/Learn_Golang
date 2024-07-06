package biz

import (
	"Learn_Golang/common"
	"Learn_Golang/modules/item/model"
	"context"
)

type ListItemStore interface {
	ListItem(ctx context.Context, filter *model.Filter, paging *common.Paging, moreKeys ...string) ([]model.TodoItem, error)
}
type listItemBiz struct {
	store ListItemStore
}

func NewListItemBiz(store ListItemStore) *listItemBiz {
	return &listItemBiz{store: store}

}

func (biz *listItemBiz) ListItem(ctx context.Context, filter *model.Filter, paging *common.Paging) ([]model.TodoItem, error) {
	data, err := biz.store.ListItem(ctx, filter, paging)

	if err != nil {
		return nil, err
	}
	return data, nil

}
