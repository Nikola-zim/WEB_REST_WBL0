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
