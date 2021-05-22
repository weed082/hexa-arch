package core

import (
	"context"
	"net/http"
	"regexp"
	"strings"
)

type key int

const (
	paramKey key = iota
)
const anyMethod = "*"

type Router struct {
	routes      map[string]map[string]http.HandlerFunc
	regexRoutes map[string]map[*regexp.Regexp]http.HandlerFunc
}

// ---------------- (1) Router 객체 생성 ----------------
func NewRouter() *Router {
	return &Router{
		routes:      make(map[string]map[string]http.HandlerFunc),
		regexRoutes: make(map[string]map[*regexp.Regexp]http.HandlerFunc),
	}
}

// ---------------- (2) Route 추가 ----------------
func (router *Router) Get(path string, handleFunc http.HandlerFunc) {
	router.addRoute("GET", path, handleFunc)
}

func (router *Router) Post(path string, handleFunc http.HandlerFunc) {
	router.addRoute("POST", path, handleFunc)
}

func (router *Router) Put(path string, handleFunc http.HandlerFunc) {
	router.addRoute("PUT", path, handleFunc)
}

func (router *Router) Delete(path string, handleFunc http.HandlerFunc) {
	router.addRoute("DELETE", path, handleFunc)
}

func (router *Router) HandleFunc(path string, handleFunc http.HandlerFunc) {
	router.addRoute(anyMethod, path, handleFunc)
}

// URL 파라미터가 필요한지에 따라 각각 regexRoutes, routes에 저장
func (router *Router) addRoute(method, path string, handleFunc http.HandlerFunc) {
	if strings.Contains(path, ":") {
		_, ok := router.regexRoutes[method]
		if !ok {
			router.regexRoutes[method] = make(map[*regexp.Regexp]http.HandlerFunc)
		}
		regexPath := router.makeRegexPath(path)
		router.regexRoutes[method][regexPath] = handleFunc
		return
	}
	_, ok := router.routes[method]
	if !ok {
		router.routes[method] = make(map[string]http.HandlerFunc)
	}
	router.routes[method][path] = handleFunc
}

func (router *Router) makeRegexPath(path string) *regexp.Regexp {
	customParamRegex := regexp.MustCompile(`:.[^(]+`)
	path = regexp.MustCompile(`:.[^/]+(?:\([^/]+\))`).ReplaceAllStringFunc(path, func(s string) string {
		s = customParamRegex.ReplaceAllStringFunc(s, func(s string) string {
			return "(?P<" + s[1:] + ">"
		})
		return s + ")"
	})
	path = regexp.MustCompile(`(:[^/]+)`).ReplaceAllStringFunc(path, func(s string) string {
		return "(?P<" + s[1:] + ">([^/]+))"
	})
	return regexp.MustCompile("^" + path + "$")
}

// ---------------- (3) 요청 리스너 ----------------
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
		params[paramKey] = paramValues[i]
	}
	context := context.WithValue(r.Context(), paramKey, params)
	return r.WithContext(context)
}

func URLParam(r *http.Request) map[string]string {
	if params := r.Context().Value(paramKey); params != nil {
		return params.(map[string]string)
	}
	return nil
}
