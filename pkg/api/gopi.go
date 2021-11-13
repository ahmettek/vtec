package gopi

import (
	"net/http"
	"strings"
)

type Gopi struct {
	routes  []Route
	mux *http.ServeMux
}

type Route struct {
	Method string
	Path   Path
	Handler   func(w http.ResponseWriter, r *http.Request)
}

type Path struct {
	params []string
	absolutePath string
	splitPath []string
}

type basicApiHandler struct {
	api *Gopi
}

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
	routes := append(e.routes, Route{Method: method, Path: parsePath(path),Handler: handler})
	e.routes=routes;
}

func parsePath(url string) Path{

	routData :=&Path{};
	absolutePath :=""
	split := strings.Split(url, "/")
	for i, s := range split {

		if strings.Contains(s, ":") {
			routData.params = append(routData.params, s)
			routData.splitPath= append(routData.splitPath, s)
		}else{
			absolutePath += s+"/"
			routData.splitPath= append(routData.splitPath, s)
		}
		println(i)
	}
	routData.absolutePath =absolutePath

	return *routData
}

func (h *basicApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	for i := range h.api.routes {

		rPath := parsePath(r.URL.Path)
		curPath := h.api.routes[i].Path

		if h.api.routes[i].Method == r.Method && curPath.absolutePath == rPath.absolutePath && len(curPath.params) == len(rPath.params) {

			success := true
			for j := range rPath.splitPath {
				if !strings.HasPrefix(curPath.splitPath[j], ":") && rPath.splitPath[j] != curPath.splitPath[j] {
				success=false
				break
				}
			}

			if success {
				h.api.routes[i].Handler(w,r)
			}

		}
	}
}
