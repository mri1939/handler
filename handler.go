// Package handler contains some simple helpers to building web using standard net/http package
package handler

import (
	"net/http"
	"strings"
	"sync"
)

// MethodHandler is a handler that handle request on specific method
type MethodHandler struct {
	mu             sync.Mutex
	handler        map[string]http.Handler
	defaultHandler http.Handler
}

// NewMethodHandler create a new MethodHandler with a defaultHandler value
func NewMethodHandler(defaultHandler http.Handler) *MethodHandler {
	return &MethodHandler{
		defaultHandler: defaultHandler,
		handler:        make(map[string]http.Handler),
	}
}

// Add handler for specified method
func (m *MethodHandler) Add(method string, handler http.Handler) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.handler[method] = handler
}

// Handle the request using handler that matched the method
// if no method could be found on the map
// serve with default handler
func (m *MethodHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.mu.Lock()
	defer m.mu.Unlock()
	key := r.Method
	if r.Method == http.MethodHead {
		key = http.MethodGet
	}
	if handler, ok := m.handler[key]; ok {
		handler.ServeHTTP(w, r)
		return
	}
	if r.Method == http.MethodOptions {
		f := func(w http.ResponseWriter, r *http.Request) {
			methods := make([]string, 0)
			for k := range m.handler {
				methods = append(methods, k)
			}
			w.Header().Add("Allow", strings.Join(methods, ", "))
		}
		http.HandlerFunc(f).ServeHTTP(w, r)
		return
	}
	m.defaultHandler.ServeHTTP(w, r)
}

// GetURIParam returns a slice of string of parameters parsed from given URI
func GetURIParam(prefix, uri string) []string {
	uri = strings.Trim(strings.TrimPrefix(uri, prefix), "/")
	if len(uri) < 1 {
		return nil
	}
	params := strings.Split(uri, "/")
	return params
}

// ResponseEncoder are interface to encodes data to any format and write the data to ResponseWriter
// Encode must not write anything to responseWriter if there is error happened.
// Instead it must return the error immediately before write anything to the ResponseWriter
type ResponseEncoder interface {
	Encode(http.ResponseWriter, interface{}) error
}

// HandleData returns a handler with data that was encoded with ResponseEncoder
// and written to the ResponseWriter
func HandleData(e ResponseEncoder, errorHandler http.Handler, data interface{}) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := e.Encode(w, data)
		if err != nil {
			errorHandler.ServeHTTP(w, r)
			return
		}
	})
}
