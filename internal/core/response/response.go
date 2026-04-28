package response

import (
	"reflect"

	"github.com/gin-gonic/gin"
)

type WebResponse struct {
	RC        string      `json:"rc"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	Errors    interface{} `json:"errors,omitempty"`
	Meta      interface{} `json:"meta,omitempty"`
	RequestID string      `json:"request_id"`
}

func SendSuccess(c *gin.Context, status Status, data interface{}, meta ...interface{}) {
	if isEmpty(data) {
		SendError(c, DataNotFound, "")
		return
	}

	requestID, _ := c.Get("request_id")
	requestIDStr, _ := requestID.(string)

	var metaData interface{}
	if len(meta) == 1 {
		metaData = meta[0]
	} else if len(meta) > 1 {
		metaData = meta
	}

	c.JSON(status.HttpCode, WebResponse{
		RC:        status.RC,
		Message:   status.Message,
		Data:      data,
		Meta:      metaData,
		RequestID: requestIDStr,
	})
}

func isEmpty(data interface{}) bool {
	if data == nil {
		return true
	}

	v := reflect.ValueOf(data)
	switch v.Kind() {
	case reflect.Slice, reflect.Map, reflect.Chan, reflect.Array:
		return v.Len() == 0
	case reflect.Ptr:
		if v.IsNil() {
			return true
		}
		return isEmpty(v.Elem().Interface())
	case reflect.Interface:
		if v.IsNil() {
			return true
		}
		return isEmpty(v.Elem().Interface())
	default:
		return false
	}
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
