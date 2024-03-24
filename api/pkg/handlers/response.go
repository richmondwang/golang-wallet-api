package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

// Response response for the API
type Response interface {
	Render(w http.ResponseWriter, r *http.Request) error
	WithCode(code int) Response
}

// ResponseWrapper response instance for the API
type ResponseWrapper struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error,omitempty"`
	Code  int         `json:"code"`
}

// WithCode sets the code for the response instance
func (r *ResponseWrapper) WithCode(code int) Response {
	r.Code = code
	return r
}

// NewResponse creates a response instance
func NewResponse(data interface{}) Response {
	return &ResponseWrapper{
		Data: data,
		Code: 200,
	}
}

// NewResponseError create an error response instance
func NewResponseError(err error) Response {
	return &ResponseWrapper{
		Error: err.Error(),
		Code:  500,
	}
}

// Render pre-render for the response instance
func (rw *ResponseWrapper) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, rw.Code)
	return nil
}
