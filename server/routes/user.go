package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/supermario64bit/whatsapp_connect/server/controller"
)

// Includes all the routes for organisation
func mountUserRoutes(r *gin.Engine) {
	userRouteGroup := r.Group("/user")
	{
		ctrl := controller.NewUserController()

		userRouteGroup.POST("", ctrl.Create)
		userRouteGroup.GET("", ctrl.Find)
		userRouteGroup.GET("/:id", ctrl.FindByID)
		userRouteGroup.PUT("/:id", ctrl.UpdateByID)
		userRouteGroup.DELETE("/:id", ctrl.DeleteByID)
	}
}
