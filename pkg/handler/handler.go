package handler

import (
	"WEB_REST_exm0302/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Static("/static/css", "./static/css")
	router.LoadHTMLGlob("static/templates/*.html")
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItem)
				items.GET("/:item_id", h.getItemById)
				items.PUT("/:item_id", h.updateItem)
				items.DELETE("/:item_id", h.deleteItem)
			}
		}

	}
	testJson := router.Group("/test_json")
	{
		testJson.GET("/readId", h.readId)
		testJson.POST("/write", h.writeJson)
		testJson.POST("/showJson", h.showJson)
	}

	test := router.Group("/test")
	{
		test.GET("/", h.showTestHome)
		test.POST("/postform", h.showResultTestHome)
	}

	return router
}
