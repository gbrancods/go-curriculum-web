package router

import (
	"net/http"

	"github.com/igab-dev/go-curriculum-web/pkg/web/handlers"
)

func HandleRequest() (err error) {
	http.HandleFunc("/", handlers.Home)
	err = http.ListenAndServe(":8082", nil)
	return
}
