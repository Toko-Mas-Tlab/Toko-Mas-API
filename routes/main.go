package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(DB *gorm.DB, r *gin.Engine) []*gin.RouterGroup {
	var routes []*gin.RouterGroup

	routes = append(routes, routeJenisBarang(DB, r))
	routes = append(routes, routeAnggota(DB, r))

	return routes
}
