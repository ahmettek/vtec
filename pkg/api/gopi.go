package api

import (
	"net/http"
)

type Gopi struct {
	routes  []Route
	mux *http.ServeMux
}

type Route struct {
	Method string `json:"method"`
	Path   string `json:"path"`
	Handler   func(w http.ResponseWriter, r *http.Request) `json:"name"`
}

type basicApiHandler struct {
	api *Gopi
}

// New returns an initialized Gopi structure, ready to use.
func New() *Gopi {
	mux := http.NewServeMux()
	return &Gopi{
		mux: mux,
	}
}

func (g*Gopi) Serve(port string) {
	g.mux.Handle("/",&basicApiHandler{g})
	http.ListenAndServe(":"+port, g.mux)
}

func (e *Gopi) GET(path string,handler func(w http.ResponseWriter, r *http.Request)) {
	 e.add(http.MethodGet,path,handler)
}

func (e *Gopi) POST(path string,handler func(w http.ResponseWriter, r *http.Request)) {
	e.add(http.MethodPost,path,handler)
}

func (e *Gopi) add(method string,path string, handler func(w http.ResponseWriter, r *http.Request)) {
	routes := append(e.routes, Route{Method: method, Path: path,Handler: handler})
	e.routes=routes;
}

func (h *basicApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")

	for i := range h.api.routes {
		if h.api.routes[i].Method == r.Method && h.api.routes[i].Path == r.URL.Path {
			h.api.routes[i].Handler(w,r)
		}
	}
}