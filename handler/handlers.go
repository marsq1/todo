package handler

import "github.com/gin-gonic/gin"

type Handler struct {

}

func NewHandler() *Handler {
	return &Handler{}
}


func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sing-up", h.singUp)
		auth.GET("/sing-in", h.singIn)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllList)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			items := lists.Group("/:id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.createItem)
				items.PUT("/:item_id", h.updateItem)
				items.GET("/:item_id", h.getItemById)
				items.DELETE("/:item_id", h.deleteItem)
			}
		}
	}
	return router
}