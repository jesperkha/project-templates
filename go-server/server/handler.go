package server

import "net/http"

func pingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ping"))
	}
}

func indexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/index.html")
	}
}

func assetsHandler() http.HandlerFunc {
	return http.StripPrefix("/assets/", http.FileServer(http.Dir("web/assets"))).ServeHTTP
}
