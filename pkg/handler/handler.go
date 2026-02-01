package handler

import (
	"github.com/egor/watcher/pkg/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}

	api := router.Group("/api", h.userIdentity)
	{

		targets := api.Group("/targets")
		{
			targets.POST("/", h.CreateTarget)
			targets.GET("/", h.GetAllTarget)
			targets.GET("/:id", h.GetTargetById)
			targets.PUT("/:id", h.UpdateTarget)
			targets.DELETE("/:id", h.DeleteTarget)

		}

	}
	return router
}
