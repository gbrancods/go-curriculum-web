package router

import (
	"net/http"

	"github.com/gbrancods/go-curriculum-web/web/handlers"
)

func HandleRequest() (err error) {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/br/", handlers.HomePtBr)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	err = http.ListenAndServe(":80", nil)
	return
}
