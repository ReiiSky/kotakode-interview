package v1

import (
	"fmt"
	"net/http"

	"github.com/Satssuki/Go-Service-Boilerplate/helpers"
	"github.com/Satssuki/Go-Service-Boilerplate/helpers/api"
	v1s "github.com/Satssuki/Go-Service-Boilerplate/services/api/v1"
	"github.com/gin-gonic/gin"
)

// InsertSingleActivity sample controller to perform insert user function
func InsertSingleActivity(c *gin.Context) {
	defer c.Request.Body.Close()

	service := v1s.CreateSingleActivityService()
	err := helpers.ReadByteAndParse(c.Request.Body, &service.Activity)

	if err == nil {
		err := service.Insert()
		if err == nil {
			c.String(201, "activity created")
			return
		}
	}

	api.JSONResponse(http.StatusBadRequest, c.Writer, gin.H{
		"message": fmt.Sprintf("error: %v", err),
	})
}

// DeleteSingleActivity sample controller to perform insert user function
func DeleteSingleActivity(c *gin.Context) {
	defer c.Request.Body.Close()

	service := v1s.CreateSingleActivityService()

	result := service.DeleteOne(c.Param("id"))
	if result > 0 {
		c.String(200, "activity deleted")
		return
	}

	api.JSONResponse(http.StatusBadRequest, c.Writer, gin.H{
		"message": "not found activity with id: " + c.Param("id"),
	})
}
