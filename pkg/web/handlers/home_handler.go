package handlers

import (
	"net/http"

	"github.com/igab-dev/go-curriculum-web/pkg/web/html"
	"github.com/igab-dev/go-curriculum-web/pkg/web/models"
)

func Home(w http.ResponseWriter, r *http.Request) {

	var p models.HomeParams
	p.Title = models.PageTitle()
	p.Name = models.Name()
	p.Age = models.Age()
	p.AboutProject = models.AboutProject()
	p.AboutMe = models.AboutMe()
	p.Experiences = models.Experiences()
	p.Educations = models.Educations()

	html.HomeParse(w, p)
}
