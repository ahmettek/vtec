package gopi

import (
	"github.com/ahmettek/vtec/pkg/logger"
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

func (g *Gopi) Listen(port string) error{
	g.mux.Handle("/", &basicApiHandler{g})
	err :=http.ListenAndServe(":"+port, g.mux)
	return err
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

func (e *Gopi) DELETE2(path string, handler func(c *GopiContext)) {
	e.add(http.MethodDelete, path, handler)
}

func (e *Gopi) HealthCheck(path string) {
	e.add(http.MethodDelete, path, func(c *GopiContext) {
		c.Res.WriteHeader(http.StatusOK)
	})
}

func (e *Gopi) add(method string, path string, handler func(c *GopiContext)) {
	routes := append(e.routes, Route{Method: method, Path: parsePath(path), Handler: handler})
	e.routes = routes
}

func parsePath(url string) Path {

	routData := &Path{}
	split := strings.Split(url, "/")

	for i := 0; i <len(split); i++ {
		if strings.Contains(split[i], ":") {
			routData.splitPath = append(routData.splitPath, split[i])
		} else {
			routData.splitPath = append(routData.splitPath, split[i])
		}
	}
	return *routData
}

func (h *basicApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	logger.Info.Printf("[%s] %s %s %s",
		r.Method,
		r.Host,
		r.URL.Path,
		r.URL.RawQuery)

	w.Header().Set("content-type", "application/json")

	context := &GopiContext{
		Param: make(map[string]string),
		Res: w,
		Req: r,
	}

	h.Handle(context)
}

func (h *basicApiHandler)  Handle(gc*GopiContext)  {
	for i := range h.api.routes {

		reqPath := parsePath(gc.Req.URL.Path)
		curPath := h.api.routes[i].Path

		if h.api.routes[i].Method == gc.Req.Method && len(reqPath.splitPath) == len(curPath.splitPath) {

			for j := range reqPath.splitPath {
				if !strings.HasPrefix(curPath.splitPath[j], ":") && reqPath.splitPath[j] != curPath.splitPath[j] {
					break
				}
			}

			gc.Param = BindParams(reqPath,curPath)

			h.api.routes[i].Handler(gc)

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
