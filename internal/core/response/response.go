package response

import (
	"github.com/gin-gonic/gin"
)

type WebResponse struct {
	RC      string      `json:"rc"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

func SendSuccess(c *gin.Context, status Status, data interface{}) {
	c.JSON(status.HttpCode, WebResponse{
		RC:      status.RC,
		Message: status.Message,
		Data:    data,
	})
}

func SendError(c *gin.Context, status Status, errDetail interface{}) {
	// If errDetail is nil or an empty string, set it to nil so omitempty works
	if errDetail == "" || errDetail == nil {
		errDetail = nil
	}

	c.JSON(status.HttpCode, WebResponse{
		RC:      status.RC,
		Message: status.Message,
		Errors:  errDetail,
	})
}
