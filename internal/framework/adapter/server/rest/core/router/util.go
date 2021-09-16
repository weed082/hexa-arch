package router

import "net/http"

// ---------------- (?) util ----------------
func UrlParam(r *http.Request) map[string]string {
	if params := r.Context().Value(paramKey); params != nil {
		return params.(map[string]string)
	}
	return nil
}
