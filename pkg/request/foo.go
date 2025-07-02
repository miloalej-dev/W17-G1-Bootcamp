package request

import (
	"errors"
	"net/http"
	"strings"
)

type FooRequest struct {
	Name        *string `json:"name"`
	Description string  `json:"description"`
}

func (p *FooRequest) Bind(r *http.Request) error {
	if p.Name == nil {
		return errors.New("name not be null")
	}

	// Post-processing after JSON decode
	*p.Name = strings.ToLower(*p.Name) // Transform data

	return nil
}
