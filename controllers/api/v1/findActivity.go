package v1

import (
	"fmt"
	"net/http"

	"github.com/Satssuki/Go-Service-Boilerplate/helpers/api"
	v1s "github.com/Satssuki/Go-Service-Boilerplate/services/api/v1"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// FindActivity sample controller to perform find user function
func FindActivity(c *gin.Context) {
	defer c.Request.Body.Close()

	service := v1s.CreateBulkActivityService()

	result, err := service.Find(bson.M{})
	if err == nil {
		api.JSONResponse(200, c.Writer, gin.H{
			"list": result,
		})
		return
	}

	api.JSONResponse(http.StatusBadRequest, c.Writer, gin.H{
		"message": fmt.Sprintf("error: %v", err),
	})
}

// ActivityDetail sample controller to perform find user function
func ActivityDetail(c *gin.Context) {
	defer c.Request.Body.Close()

	service := v1s.CreateSingleActivityService()

	result, err := service.FindByID(c.Param("id"))
	if err == nil {
		api.JSONResponse(200, c.Writer, gin.H{"result": result.Activity})
		return
	}

	api.JSONResponse(http.StatusBadRequest, c.Writer, gin.H{
		"message": fmt.Sprintf("error: %v", err),
	})
}
