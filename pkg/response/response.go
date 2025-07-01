package response

import (
	"github.com/go-chi/render"
	"net/http"
)

func SetResponse(w http.ResponseWriter, r *http.Request, status int, responseBody any) {
	render.Status(r, status)
	render.JSON(w, r, responseBody)
}
