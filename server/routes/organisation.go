package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/supermario64bit/whatsapp_connect/server/controller"
)

// Includes all the routes for organisation
func mountOrganisationRoutes(r *gin.Engine) {
	orgRouteGroup := r.Group("/organisation")
	{
		ctrl := controller.NewOrganisationController()

		orgRouteGroup.POST("", ctrl.Create)
		orgRouteGroup.GET("", ctrl.Find)
		orgRouteGroup.GET("/:id", ctrl.FindByID)
		orgRouteGroup.PUT("/:id", ctrl.UpdateByID)
		orgRouteGroup.DELETE("/:id", ctrl.DeleteByID)
	}
}
