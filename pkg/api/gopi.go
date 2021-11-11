package api

import (
	"net/http"
)

type Gopi struct {
	routes  []Route
	mux *http.ServeMux
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
	 e.Add(http.MethodGet,path,handler)
}

func (e *Gopi) Add(method string,path string, handler func(w http.ResponseWriter, r *http.Request)) {
	routes := append(e.routes, Route{Method: method, Path: path,Handler: handler})
	e.routes=routes;
}

type Route struct {
	Method string `json:"method"`
	Path   string `json:"path"`
	Handler   func(w http.ResponseWriter, r *http.Request) `json:"name"`
}

type basicApiHandler struct {
	api *Gopi
}

func (h *basicApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")

	for i := range h.api.routes {
		if h.api.routes[i].Method == r.Method && h.api.routes[i].Path == r.URL.Path {
			h.api.routes[i].Handler(w,r)
		}
	}

}