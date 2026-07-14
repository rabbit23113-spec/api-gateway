package ports

import "net/http"

type Server interface {
	Serve(port string, handler http.Handler) error
	Shutdown() error
}
