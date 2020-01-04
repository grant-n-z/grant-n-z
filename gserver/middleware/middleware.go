package middleware

import "net/http"

type middleware func(http.HandlerFunc) http.HandlerFunc

type Middleware struct {
	middlewares []middleware
}

// Constructor
func NewMiddleware(mws ...middleware) Middleware {
	return Middleware{append([]middleware(nil), mws...)}
}

// Middleware and handler process
func (m Middleware) handler(h http.HandlerFunc) http.HandlerFunc {
	for i := range m.middlewares {
		h = m.middlewares[len(m.middlewares)-1-i](h)
	}
	return h
}
