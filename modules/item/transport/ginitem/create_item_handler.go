package ginitem

import (
	"Learn_Golang/common"
	"Learn_Golang/modules/item/biz"
	"Learn_Golang/modules/item/model"
	"Learn_Golang/modules/item/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func CreateItem(db *gorm.DB) func(*gin.Context) {
	return func(context *gin.Context) {
		var data model.TodoItemCreation
		if err := context.ShouldBind(&data); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		business := biz.NewCreateItemBiz(storage.NewSQLStore(db))
		if err := business.CreateItem(context.Request.Context(), &data); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
