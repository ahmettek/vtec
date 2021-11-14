package gopi

import (
	"net/http"
	"strings"
)

type Gopi struct {
	routes []Route
	mux    *http.ServeMux
}

type GopiContext struct {
	Param map[string]string
	Res   http.ResponseWriter
	Req   *http.Request
}

type Route struct {
	Method  string
	Path    Path
	Handler func(c *GopiContext)
}

type Path struct {
	params       []string
	absolutePath string
	splitPath    []string
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

func (g *Gopi) Serve(port string) {
	g.mux.Handle("/", &basicApiHandler{g})
	http.ListenAndServe(":"+port, g.mux)
}

func (e *Gopi) GET(path string, handler func(c *GopiContext)) {
	e.add(http.MethodGet, path, handler)
}

func (e *Gopi) POST(path string, handler func(c *GopiContext)) {
	e.add(http.MethodPost, path, handler)
}

func (e *Gopi) DELETE(path string, handler func(c *GopiContext)) {
	e.add(http.MethodDelete, path, handler)
}

func (e *Gopi) add(method string, path string, handler func(c *GopiContext)) {
	routes := append(e.routes, Route{Method: method, Path: parsePath(path), Handler: handler})
	e.routes = routes
}

func parsePath(url string) Path {

	routData := &Path{}
	absolutePath := ""
	split := strings.Split(url, "/")
	for i, s := range split {

		if strings.Contains(s, ":") {
			routData.params = append(routData.params, s)
			routData.splitPath = append(routData.splitPath, s)
		} else {
			absolutePath += s + "/"
			routData.splitPath = append(routData.splitPath, s)
		}
		println(i)
	}
	routData.absolutePath = absolutePath

	return *routData
}

func (h *basicApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	context := &GopiContext{Param: make(map[string]string), Res: w, Req: r}
	for i := range h.api.routes {

		rPath := parsePath(r.URL.Path)
		curPath := h.api.routes[i].Path

		if h.api.routes[i].Method == r.Method && len(rPath.splitPath) == len(curPath.splitPath) {

			for j := range rPath.splitPath {
				if !strings.HasPrefix(curPath.splitPath[j], ":") && rPath.splitPath[j] != curPath.splitPath[j] {
					break
				}
			}

			context.Param = BindParams(rPath,curPath)

			h.api.routes[i].Handler(context)
			break
		}
	}
}
func BindParams(reqPath Path, curPath Path) map[string]string {
	params := make(map[string]string)
	for j := range reqPath.splitPath {
		if strings.HasPrefix(curPath.splitPath[j], ":") {
			params[curPath.splitPath[j]] = reqPath.splitPath[j]
		}
	}
	return params
}

func InitContext(rPath Path, curPath Path, w http.ResponseWriter, r *http.Request) *GopiContext {
	context := &GopiContext{Param: make(map[string]string), Res: w, Req: r}
	for j := range rPath.splitPath {
		if strings.HasPrefix(curPath.splitPath[j], ":") {
			context.Param[curPath.splitPath[j]] = rPath.splitPath[j]
		}
	}

	return context
}
