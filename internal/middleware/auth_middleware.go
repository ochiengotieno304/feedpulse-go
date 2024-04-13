package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ochiengotieno304/feedpulse-go/configs"
)

func ProfilerMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Get("")
		begin := time.Now().UnixMilli()
		next.ServeHTTP(w, r)
		end := time.Now().UnixMilli()
		ns := end - begin
		fmt.Printf("Request is processed in %d milliseconds\n", ns)
	}

	return http.HandlerFunc(fn)
}

func RapidProxySecretCheck(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		proxySecret := r.Header.Get("HTTP_X_MASHAPE_PROXY_SECRET")
		configs, err := configs.LoadConfig()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		secret := configs.RapidApiProySecret

		if proxySecret == "" && secret != proxySecret {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Forbidden"))
			return
		}
		next.ServeHTTP(w, r)

		fmt.Println("Check Complete")
	}

	return http.HandlerFunc(fn)
}

// HTTP_X_MASHAPE_PROXY_SECRET
