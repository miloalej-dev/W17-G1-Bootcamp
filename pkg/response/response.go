package response

import (
	"github.com/go-chi/render"
	"net/http"
)

type Response struct {
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	StatusCode int         `json:"-"`
}

func (re *Response) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before marshalling to JSON
	render.Status(r, re.StatusCode)
	return nil
}

func NewResponse(data interface{}, statusCode int) *Response {
	resp := &Response{
		Message:    "",
		Data:       data,
		StatusCode: statusCode,
	}
	return resp
}

func NewErrorResponse(message string, statusCode int) *Response {
	resp := &Response{
		Message:    message,
		StatusCode: statusCode,
	}
	return resp
}
