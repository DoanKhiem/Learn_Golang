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

func ListItem(db *gorm.DB) func(*gin.Context) {
	return func(context *gin.Context) {
		var paging common.Paging
		if err := context.ShouldBind(&paging); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		paging.Process()
		var filter model.Filter

		if err := context.ShouldBind(&filter); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := storage.NewSQLStore(db)
		business := biz.NewListItemBiz(store)

		result, err := business.ListItem(context.Request.Context(), &filter, &paging)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
