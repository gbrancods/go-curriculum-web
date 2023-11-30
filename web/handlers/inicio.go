package handlers

import (
	"net/http"

	"github.com/gbrancods/go-curriculum-web/web/html"
	"github.com/gbrancods/go-curriculum-web/web/models"
)

func Início(w http.ResponseWriter, r *http.Request) {

	p := models.Home("pt")

	html.HomeParse(w, p)
}
