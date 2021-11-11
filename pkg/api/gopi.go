package api

import (
	"encoding/json"
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
	g.mux.Handle("/api/keys",&basicApiHandler{})
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

type basicApiHandler struct {}

func (h *basicApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodGet:
		h.Get(w, r)
		return
	case r.Method == http.MethodPut:
		h.Set(w, r)
		return
	default:
		notFound(w, r)
		return
	}
}

func (h *basicApiHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(true)
}

func (h *basicApiHandler) Set(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(true)
}
func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
}