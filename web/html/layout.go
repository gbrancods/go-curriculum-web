package html

import (
	"embed"
	"fmt"
	"net/http"
	"text/template"

	"github.com/gbrancods/go-curriculum-web/web/models"
)

// Necessary keep this on same folder of html files
// The comment go:embed is necessary to set folder

//go:embed *
var files embed.FS

func parse(file string) *template.Template {
	return template.Must(
		template.New("layout.html").ParseFS(files, "layout.html", file))
}

func HomeParse(w http.ResponseWriter, p models.HomeParams) {
	home := parse("home.html")
	if err := home.Execute(w, p); err != nil {
		fmt.Println("Error parsing html", err)
	}
}
