package router

import (
	"errors"

	"github.com/iamvineettiwari/go-web-server/http"
)

type Router struct {
	methods map[string]*Tree
}

func NewRouter() *Router {
	return &Router{
		methods: make(map[string]*Tree),
	}
}

func (r *Router) registerHandler(method string, path string, handler func(req *http.Request, res *http.Response)) {
	if _, present := r.methods[method]; !present {
		r.methods[method] = NewTree()
	}

	r.methods[method].Insert(path, handler)
}

func (r *Router) Get(path string, handler http.HandlerFunc) {
	r.registerHandler(http.GET, path, handler)
}

func (r *Router) Post(path string, handler http.HandlerFunc) {
	r.registerHandler(http.POST, path, handler)
}

func (r *Router) Put(path string, handler http.HandlerFunc) {
	r.registerHandler(http.PUT, path, handler)
}

func (r *Router) Delete(path string, handler http.HandlerFunc) {
	r.registerHandler(http.DELETE, path, handler)
}

func (r *Router) Resolve(method string, path string) (http.HandlerFunc, http.Params, error) {
	if _, present := r.methods[method]; !present {
		return nil, nil, errors.New("Not found")
	}

	handlerFunc, params := r.methods[method].Resolve(path)

	if handlerFunc == nil {
		return nil, nil, errors.New("Not found")
	}

	return handlerFunc, params, nil
}
