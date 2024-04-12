package response

import "github.com/gofiber/fiber/v2"

type BodyStruct struct {
	StatusCode int         `json:"statusCode,omitempty"`
	Message    string      `json:"message,omitempty"`
	Err        interface{} `json:"error,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

// error sets the error information in the BodyStruct
func (b *BodyStruct) Error(statusCode int, msg string, err interface{}, data interface{}) *BodyStruct {
	b.StatusCode = statusCode
	b.Message = msg
	b.Err = err
	b.Data = data

	return b
}

// success sets the success information
func (b *BodyStruct) Success(statuscode int, msg string, data interface{}) *BodyStruct {
	b.StatusCode = statuscode
	b.Message = msg
	b.Data = data

	return b
}

// used to send response with fiber ctx
func (b *BodyStruct) Send(ctx *fiber.Ctx) error {
	return ctx.JSON(b)
}
