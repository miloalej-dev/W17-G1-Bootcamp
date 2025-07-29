package response

import (
	"github.com/go-chi/render"
	"net/http"
)

type Response struct {
	//Data	render.Renderer	`json:"data"`
	Data       any    `json:"data,omitempty"`
	Message    string `json:"message,omitempty"`
	StatusCode int    `json:"-"`
}

func (re *Response) Render(w http.ResponseWriter, r *http.Request) error {
	// Set proper UTF-8 Content-Type header
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// Pre-processing before marshalling to JSON
	render.Status(r, re.StatusCode)
	return nil
}

func NewResponse(data any, statusCode int) *Response {
	resp := &Response{
		Data:       data,
		StatusCode: statusCode,
	}
	return resp
}

func NewErrorResponse(msg string, statusCode int) *Response {
	resp := &Response{
		Message:    msg,
		StatusCode: statusCode,
	}
	return resp
}
