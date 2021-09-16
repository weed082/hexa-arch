package router

import (
	"context"
	"net/http"
	"regexp"
	"strings"
)

type key int
type middleware func(http.Handler) http.Handler

const paramKey key = iota

type Router struct {
	routes      map[string]map[string]http.HandlerFunc
	regexRoutes map[string]map[*regexp.Regexp]http.HandlerFunc
	middlewares []middleware
}

// ---------------- (1) Router 객체 생성 ----------------
func New() *Router {
	return &Router{
		routes:      make(map[string]map[string]http.HandlerFunc),
		regexRoutes: make(map[string]map[*regexp.Regexp]http.HandlerFunc),
	}
}

// ---------------- (2) Route 추가 ----------------
func (router *Router) Get(path string, handleFunc http.HandlerFunc) {
	router.addRoute(http.MethodGet, path, handleFunc)
}

func (router *Router) Post(path string, handleFunc http.HandlerFunc) {
	router.addRoute(http.MethodPost, path, handleFunc)
}

func (router *Router) Put(path string, handleFunc http.HandlerFunc) {
	router.addRoute(http.MethodPut, path, handleFunc)
}

func (router *Router) Delete(path string, handleFunc http.HandlerFunc) {
	router.addRoute(http.MethodDelete, path, handleFunc)
}

func (router *Router) addRoute(method, path string, handleFunc http.HandlerFunc) {
	if strings.Contains(path, ":") {
		_, ok := router.regexRoutes[method]
		if !ok {
			router.regexRoutes[method] = make(map[*regexp.Regexp]http.HandlerFunc)
		}
		path = regexp.MustCompile(`:.[^/]+(?:\([^/]+\))`).ReplaceAllStringFunc(path, router.makeCustomRegexParam)
		path = regexp.MustCompile(`(:[^/]+)`).ReplaceAllStringFunc(path, router.makeRegexParam)
		regexPath := regexp.MustCompile("^" + path + "$")
		chainedHandler := router.chain(handleFunc)
		router.regexRoutes[method][regexPath] = chainedHandler.ServeHTTP
		return
	}
	_, ok := router.routes[method]
	if !ok {
		router.routes[method] = make(map[string]http.HandlerFunc)
	}
	chainedHandler := router.chain(handleFunc)
	router.routes[method][path] = chainedHandler.ServeHTTP
}

func (router *Router) makeCustomRegexParam(param string) string {
	splits := strings.SplitN(param, "(", 2)
	return "(?P<" + splits[0][1:] + ">(" + splits[1] + ")"
}

func (router *Router) makeRegexParam(param string) string {
	return "(?P<" + param[1:] + ">([^/]+))"
}

// ---------------- (3) middleware ----------------
func (router *Router) Use(middlewares ...middleware) {
	if router.middlewares == nil {
		router.middlewares = []middleware{}
	}
	router.middlewares = append(router.middlewares, middlewares...)
}

func (router *Router) chain(endpoint http.Handler) http.Handler {
	if router.middlewares == nil {
		return endpoint
	}
	handler := router.middlewares[0](endpoint)

	for i := 1; i < len(router.middlewares); i++ {
		handler = router.middlewares[i](handler)
	}
	return handler
}

// ---------------- (4) request listener ----------------
func (router *Router) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	handleFunc, ok := router.routes[r.Method][r.URL.Path]
	if ok {
		handleFunc(rw, r)
		return
	}
	for regexPath, handleFunc := range router.regexRoutes[r.Method] {
		if regexPath.MatchString(r.URL.Path) {
			req := router.makeContextWithParams(r, regexPath)
			handleFunc(rw, req)
			return
		}
	}
	http.NotFound(rw, r)
}

func (router *Router) makeContextWithParams(r *http.Request, regexPath *regexp.Regexp) *http.Request {
	paramValues := regexPath.FindStringSubmatch(r.URL.Path)
	params := make(map[string]string)

	for i, paramKey := range regexPath.SubexpNames() {
		if paramKey != "" {
			params[paramKey] = paramValues[i]
		}
	}
	context := context.WithValue(r.Context(), paramKey, params)
	return r.WithContext(context)
}
