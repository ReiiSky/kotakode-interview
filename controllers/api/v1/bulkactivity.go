package v1

import (
	"fmt"
	"net/http"

	"github.com/Satssuki/Go-Service-Boilerplate/helpers"
	"github.com/Satssuki/Go-Service-Boilerplate/helpers/api"
	v1s "github.com/Satssuki/Go-Service-Boilerplate/services/api/v1"
	"github.com/gin-gonic/gin"
)

// InsertBulkActivity sample controller to perform insert user function
func InsertBulkActivity(c *gin.Context) {
	defer c.Request.Body.Close()

	service := v1s.CreateBulkActivityService()
	err := helpers.ReadByteAndParse(c.Request.Body, &service)
	
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
