package responses

import (
	"github.com/gin-gonic/gin"
)

type CustomResponse struct {
	Message string      `json:"message"`
	Error   bool        `json:"error"`
	Data    interface{} `json:"data,omitempty"` // Use omitempty to skip this field if it's nil
}

// SendResponse is a helper function to send custom responses
func SendResponse(ctx *gin.Context, statusCode int, message string, data interface{}, isError bool) {
	response := CustomResponse{
		Message: message,
		Error:   isError,
		Data:    data,
	}
	ctx.JSON(statusCode, response)
}
