package response

import "golang.org/x/text/message"

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, Response{
		Success: true,
		Message: mmessage,
		Data:data,
	})
}

func Error (c *gin.Context,code int,message string) {
	c.JSON(code,Response{
		Success:false,
		Message:message,
	})
}