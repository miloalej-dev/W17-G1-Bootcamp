package response

import (
	"net/http"
	"github.com/go-chi/render"
)

func SetResponse(w http.ResponseWriter, r *http.Request, status int, responseBody any, err error) {
	var resp any
	if err != nil {
		resp = Error{
			Message: err.Error(),
		}
	} else {
		resp = Success{
			Data: responseBody,
		}
	}
	render.Status(r, status)
	render.JSON(w, r, resp)
}

