package response

import (
	"github.com/gin-gonic/gin"
)

type WebResponse struct {
	RC        string      `json:"rc"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	Errors    interface{} `json:"errors,omitempty"`
	RequestID string      `json:"request_id"`
}

func SendSuccess(c *gin.Context, status Status, data interface{}) {
	requestID, _ := c.Get("request_id")
	requestIDStr, _ := requestID.(string)

	c.JSON(status.HttpCode, WebResponse{
		RC:        status.RC,
		Message:   status.Message,
		Data:      data,
		RequestID: requestIDStr,
	})
}

func SendError(c *gin.Context, status Status, errDetail interface{}) {
	// If errDetail is nil or an empty string, set it to nil so omitempty works
	if errDetail == "" || errDetail == nil {
		errDetail = nil
	}

	requestID, _ := c.Get("request_id")
	requestIDStr, _ := requestID.(string)

	c.JSON(status.HttpCode, WebResponse{
		RC:        status.RC,
		Message:   status.Message,
		Errors:    errDetail,
		RequestID: requestIDStr,
	})
}
