package models

type HomeParams struct {
	Title        string
	Name         string
	Age          string
	AboutProject string
	AboutMe      []string
	Educations   []Education
	Experiences  []Experience
}

type Education struct {
	BeginAt    string
	EndAt      string
	SchoolName string
	Matery     string
}

type Experience struct {
	BeginAt      string
	EndAt        string
	BusinessName string
	Vocation     string
}

func PageTitle() (pgTitle string) {
	pgTitle = "Go Curriculum Web"
	return
}

func Name() (name string) {
	name = "Gabriel Branco da Silva"
	return
}

func Age() (name string) {
	name = "23 years old"
	return
}

func Educations() (educations []Education) {

	var ifsp Education
	ifsp.BeginAt = "2014"
	ifsp.EndAt = "2016"
	ifsp.SchoolName = "IFSP"
	ifsp.Matery = "Technical course integrated into high school - computer maintenance and support"

	var ads Education
	ads.BeginAt = "2017"
	ads.EndAt = "2017"
	ads.SchoolName = "IFSP"
	ads.Matery = "Icomplete, systems analysis and development"

	var uni Education
	uni.BeginAt = "2018"
	uni.EndAt = "2019"
	uni.SchoolName = "Unifev"
	uni.Matery = "Incomplete, advertising and marketing"

	var alu Education
	alu.BeginAt = "2020"
	alu.EndAt = "2022"
	alu.SchoolName = "Alura"
	alu.Matery = "Docker, Linux, Go, UI-UX, JavaScript, Owasp Security"

	educations = []Education{
		ifsp,
		ads,
		uni,
		alu,
	}

	return
}

func Experiences() (experiences []Experience) {

	var silk Experience
	silk.BeginAt = "jan 2016"
	silk.EndAt = "oct 2017"
	silk.Vocation = "Final Art"
	silk.BusinessName = "Silk Cromia"

	var flash Experience
	flash.BeginAt = "oct 2017"
	flash.EndAt = "may 2019"
	flash.Vocation = "Technical support"
	flash.BusinessName = "Flash Net Brasil"

	var rosa Experience
	rosa.BeginAt = "oct 2019"
	rosa.EndAt = "feb 2020"
	rosa.Vocation = "T.I. Analyst"
	rosa.BusinessName = "Rosa Mística"

	var cofer Experience
	cofer.BeginAt = "feb 2020 "
	cofer.EndAt = "apr 2021"
	cofer.Vocation = "T.I. Assistant"
	cofer.BusinessName = "COFERPOL"

	var cofdev Experience
	cofdev.BeginAt = "may 2021"
	cofdev.EndAt = "feb 2023"
	cofdev.Vocation = "Software Enginner"
	cofdev.BusinessName = "COFERPOL"

	var auto Experience
	auto.BeginAt = "feb 2023"
	auto.EndAt = "at the time"
	auto.Vocation = "Print on demmand"
	auto.BusinessName = "e-Commerce"

	experiences = []Experience{
		silk,
		flash,
		rosa,
		cofer,
		cofdev,
		auto,
	}

	return
}

func AboutMe() (about []string) {
	about = []string{
		"* ",
		"* Linux enthusiast since 2014",
	}
	return
}

func AboutProject() (abProj string) {
	abProj = "A simple Go html template project, theme based an old Go playground, you can see this on <a href='https://github.com/igab-dev/go-curriculum-web' target='_blank'>github.com/igab-dev/go-curriculum-web</a>"
	return
}
