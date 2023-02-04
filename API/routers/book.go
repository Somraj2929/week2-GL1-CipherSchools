package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/somraj2929/week2-GL1-CipherSchools/database"
	"github.com/somraj2929/week2-GL1-CipherSchools/handler"
)

func Router() *gin.Engine{
	router := gin.Default()
	api := handler.Handler {
		DB:database.GetDB(),
	}
	router.GET("/book", api.GetBooks)
	router.POST("/book", api.SaveBook)
	return router
}