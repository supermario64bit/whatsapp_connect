package routes

import "github.com/gin-gonic/gin"

// Register all http routes
func MountHTTPRoutes(r *gin.Engine) {
	mountOrganisationRoutes(r)
}
