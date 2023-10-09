package router

import (
	"net/http"

	"github.com/igab-dev/go-curriculum-web/pkg/web/handlers"
)

func HandleRequest() (err error) {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		// do NOT do this. (see below)
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	err = http.ListenAndServe(":80", nil)
	return
}
