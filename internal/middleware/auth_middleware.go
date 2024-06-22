package middleware

import (
	"fmt"
	"net/http"

	"github.com/ochiengotieno304/feedpulse-go/configs"
)

func RapidProxySecretCheck(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		proxySecret := r.Header.Get("X-Mashape-Proxy-Secret")
		configs, err := configs.LoadConfig()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		secret := configs.RapidApiProySecret

		if proxySecret == "" || secret != proxySecret {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Forbidden"))
			return
		}
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
