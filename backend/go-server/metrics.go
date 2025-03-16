package main

import (
	"fmt"
	"net/http"
)

func (cfg *config) handlerMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`
	<html>
		<body>
			<h1>Welcome Admin</h1>
			<p>%d requests have been made since start.</p>
		</body>
	</html>
	`, cfg.fileserverHits.Load())))
}

func (cfg *config) middleWareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		trackedPaths := map[string]bool{
			"/generate-resume": true,
		}

		for path := range trackedPaths {
			if len(r.URL.Path) >= len(path) && r.URL.Path[:len(path)] == path {
				cfg.fileserverHits.Add(1)
				break
			}
		}

		next.ServeHTTP(w, r)
	})
}
