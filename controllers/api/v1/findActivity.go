package v1

import (
	"fmt"
	"net/http"
	"strconv"

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

// FindByGroup sample controller to perform find user function
func FindByGroup(c *gin.Context) {
	service := v1s.CreateBulkActivityService()

	result, err := service.Find(
		bson.M{
			"group": c.Param("groupName"),
		},
	)
	if err == nil {
		api.JSONResponse(200, c.Writer, gin.H{"result": result.Activities})
		return
	}

	api.JSONResponse(http.StatusBadRequest, c.Writer, gin.H{
		"message": fmt.Sprintf("error: %v", err),
	})
}

// FindByPriority sample controller to perform find user function
func FindByPriority(c *gin.Context) {
	service := v1s.CreateBulkActivityService()
	number, err := strconv.Atoi(c.Param("priorityNumber"))
	if err == nil {
		fmt.Println(number)
		result, err := service.Find(
			bson.M{
				"priority": number,
			},
		)
		if err == nil {
			api.JSONResponse(200, c.Writer, gin.H{"result": result.Activities})
			return
		}
	}

	api.JSONResponse(http.StatusBadRequest, c.Writer, gin.H{
		"message": fmt.Sprintf("error: %v", err),
	})
}
