package middlewares

import (
	"net/http"
	"strings"
)

type CorsConfig struct {
	AllowedOrigins []string
	AllowedMethods []string
	AllowedHeaders []string
}

type CorsHandler struct {
	Next   http.Handler
	Config *CorsConfig
}

func (ch *CorsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if len(ch.Config.AllowedOrigins) > 0 {
		w.Header().Add("Access-Control-Allow-Origin", strings.Join(ch.Config.AllowedOrigins, ","))
	}
	if len(ch.Config.AllowedHeaders) > 0 {
		w.Header().Add("Access-Control-Allow-Headers", strings.Join(ch.Config.AllowedHeaders, ","))
	}
	if len(ch.Config.AllowedMethods) > 0 {
		w.Header().Add("Access-Control-Allow-Methods", strings.Join(ch.Config.AllowedMethods, ","))
	}
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	ch.Next.ServeHTTP(w, r)
}
