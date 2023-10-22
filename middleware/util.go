package middleware

import (
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type Middleware func(http.Handler) http.Handler

func ApplyMiddleware(h http.Handler, middleware ...Middleware) http.Handler {
	for i := len(middleware) - 1; i >= 0; i-- {
		h = middleware[i](h)
	}
	return h
}

type Endpoint struct {
	Path       string
	Middleware []Middleware
	Handler    http.Handler
}

type CustomMux struct {
	Mux      *runtime.ServeMux
	Endpoint []Endpoint
}

func (cm *CustomMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	for _, endpoint := range cm.Endpoint {
		if strings.HasPrefix(path, endpoint.Path) {
			wrappedHandler := ApplyMiddleware(endpoint.Handler, endpoint.Middleware...)
			wrappedHandler.ServeHTTP(w, r)
			return
		}
	}
	cm.Mux.ServeHTTP(w, r) // No middleware for other paths
}
