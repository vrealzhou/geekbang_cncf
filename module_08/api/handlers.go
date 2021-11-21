package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/vrealzhou/geekbang_cncf/module_08/log"
)

type Handler struct {
	version string
	log     *log.Logger
}

func NewHandler(version string, logger *log.Logger) *Handler {
	return &Handler{
		version: version,
		log:     logger,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	respHeaders := w.Header()
	respHeaders.Set("X-Request-From", r.RemoteAddr)
	for k, v := range r.Header {
		if strings.ToLower(k) == "x-forwarded-for" {
			respHeaders.Set("X-RequestFrom", strings.TrimSpace(strings.Split(v[0], ",")[0]))
		} else {
			respHeaders[k] = v
		}
	}
	respHeaders.Set("VERSION", h.version)

	httpStaus := http.StatusOK
	if r.URL.Path != "/healthz" {
		httpStaus = http.StatusNotFound
	}
	w.WriteHeader(httpStaus)
	// h.log.Infof("Request from %s, HTTP Code %d", r.RemoteAddr, httpStaus)
	w.Write([]byte(fmt.Sprintf(`{"status":"%s"}`, http.StatusText(httpStaus))))
}
