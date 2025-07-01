package route

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

// DefaultRoutes sets up the default routes for the API, such as the root endpoint and error handling for not found
// and method not allowed requests.
func DefaultRoutes(router chi.Router) {
	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{
			"Message": "Resource not found, please check the URL",
		})
	})

	router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		render.Status(r, http.StatusMethodNotAllowed)
		render.JSON(w, r, map[string]string{
			"Message": "Method not allowed, please check the request method",
		})
	})
}
