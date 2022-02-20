package router

import "net/http"

// ---------------- (?) util ----------------
func UrlParam(r *http.Request) map[string]string {
	params := r.Context().Value(paramKey)
	if params == nil {
		return nil
	}
	return params.(map[string]string)
}
