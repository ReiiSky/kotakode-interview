package route

import (
	v1 "github.com/Satssuki/Go-Service-Boilerplate/controllers/api/v1"
	"github.com/Satssuki/Go-Service-Boilerplate/middlewares"
	"github.com/gin-gonic/gin"
)

// SetupRouter Register each route here
func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.AppliAllCORS)

	apig := router.Group("/api")
	{
		v1g := apig.Group("/v1")
		{
			v1g.POST("/user", v1.InsertUser)
			v1g.POST("/activities", v1.InsertSingleActivity)
			v1g.POST("/activities/bulk", v1.InsertBulkActivity)

			v1g.GET("/activities", v1.FindActivity)
			v1g.GET("/activities/detail/:id", v1.ActivityDetail)
			v1g.GET("/activities/group/:groupName", v1.FindByGroup)
			v1g.GET("/activities/priority/:priorityNumber", v1.FindByPriority)
			v1g.GET("/activities/group", v1.GetGroup)

			v1g.DELETE("/activities/:id", v1.DeleteSingleActivity)
		}
	}
	return router
}
