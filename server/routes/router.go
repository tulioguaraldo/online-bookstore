package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/tulioguaraldo/online-bookstore/core/book"
	"github.com/tulioguaraldo/online-bookstore/db"
	"github.com/tulioguaraldo/online-bookstore/utils"
)

func GetRoutes() *gin.Engine {
	router := gin.Default()

	// dependencies
	repository := book.NewBookRepository(db.ConnectToDb())
	service := book.NewBookService(&repository)
	controller := book.NewBookController(&service)

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			book := v1.Group("/book")
			{
				book.GET("", controller.GetAllBooks)
				book.GET(":id", controller.GetBookByData)
				book.POST("", controller.InsertBook)
			}
		}
	}

	api.GET("/health", utils.Health)

	return router
}
