package models

func Home(lang string) (h HomeParams) {

	h.Title = "Go Curriculum Web"
	h.Name = "Gabriel Branco da Silva"

	h.AboutMe = AboutMe(lang)
	h.AboutProject = AboutProject(lang)
	h.SwitchLang = SwitchLang(lang)
	h.ExperienceTitle = ExperienceTitle(lang)
	h.Experiences = Experiences(lang)
	h.EducationTitle = EducationTitle(lang)
	h.Educations = Educations(lang)
	h.EventsTitle = EventsTitle(lang)
	h.Events = Events(lang)
	h.ArticlesTitle = ArticlesTitle(lang)
	h.Articles = Articles(lang)

	return
}

type HomeParams struct {
	Title           string
	Name            string
	SwitchLang      string
	AboutProject    string
	AboutMe         string
	EducationTitle  string
	Educations      []Education
	ExperienceTitle string
	Experiences     []Experience
	EventsTitle     string
	Events          []Event
	ArticlesTitle   string
	Articles        []Article
}

type Education struct {
	KnowledgeObtained string
	BeginAt           string
	EndAt             string
	Duration          string
	SchoolName        string
	Matery            string
	Description       string
}

type Experience struct {
	KnowledgeObtained string
	BeginAt           string
	EndAt             string
	Duration          string
	BusinessName      string
	Vocation          string
	Description       string
}

type Event struct {
	Name string
	Year int
	Type string
}

type Article struct {
	Name    string
	Year    int
	Type    string
	Authors string
}

func ExperienceTitle(l string) (m string) {
	if l == "en" {
		m = "Experiences"
		return
	}
	m = "Experiências"
	return
}

func EducationTitle(l string) (m string) {
	if l == "en" {
		m = "Educations"
		return
	}
	m = "Educação"
	return
}

func EventsTitle(l string) (m string) {
	if l == "en" {
		m = "Events"
		return
	}
	m = "Eventos"
	return
}

func ArticlesTitle(l string) (m string) {
	if l == "en" {
		m = "Articles"
		return
	}
	m = "Artigos"
	return
}

func SwitchLang(l string) (m string) {
	if l == "en" {
		m = "<a href='/br/'>Leia em Português"
		return
	}
	m = "<a href='/'>See in english"
	return
}

func AboutMe(l string) (m string) {

	if l == "en" {
		m = "<p class='subtitle'>About Me</p>" +
			"<p class='description'>" +
			"A technology lover, five and a half years of experience, the last two modeling and developing " +
			"software in Go, main knowledge in APIs, experience with Docker, GCP VMs, ec2 instances, route 53, " +
			"Gitlab CI/CD, Postgres, Nginx, Cloudflare, Firebase, Zabbix, Node.js, tested and understands the " +
			"concept of queues and Kubernetes." +
			"</p>"
		return
	}
	m = "<p class='subtitle'>Sobre mim</p>" +
		"<p class='description'>" +
		"Apaixonado por tecnologia, cinco anos e meio de experiência na área, os dois ultimos " +
		"modelando e desenvolvendo software em Go, conhecimento principal em APIs, " +
		"experiência com Docker, GCP VMs, instâncias ec2, route 53, Gitlab CI/CD, Postgres, Nginx, " +
		"Cloudflare, Firebase, Zabbix, Node.js, testou e entende o conceito de filas e Kubernetes." +
		"</p>"

	return
}

func AboutProject(l string) (p string) {

	if l == "en" {
		p = "<p class='description'>A simple Go html template project, theme based an old Go playground, you can see this on: " +
			"<a href='https://github.com/gbrancods/go-curriculum-web' target='_blank'>github.com/gbrancods/go-curriculum-web</a> </p>"
		return
	}
	p = "<p class='description'>Um simples projeto em Go utilizando html template, tema baseado no antigo Go Playground, veja este projeto em: " +
		"<a href='https://github.com/gbrancods/go-curriculum-web' target='_blank'>github.com/gbrancods/go-curriculum-web</a> </p>"

	return
}

func Educations(l string) (e []Education) {

	if l == "en" {

		var ifsp Education
		ifsp.BeginAt = "2014"
		ifsp.EndAt = "2016"
		ifsp.SchoolName = "IFSP"
		ifsp.Matery = "Technical course integrated into high school - Computer Maintenance and Support"
		ifsp.Duration = "3 years"
		ifsp.Description = "Integrated high school and technical education, UML, database, programming logic, Linux servers."

		var ads Education
		ads.BeginAt = "2017"
		ads.EndAt = "2017"
		ads.SchoolName = "IFSP"
		ads.Matery = "Systems Analysis and Development - Incomplete"
		ads.Duration = "1 year"
		ads.Description = "Introduction to statistics, reinforcement of subjects seen during technical education."

		var uni Education
		uni.BeginAt = "2018"
		uni.EndAt = "2019"
		uni.SchoolName = "Unifev"
		uni.Matery = "Advertising and Marketing - Incomplete"
		uni.Duration = "2 years"
		uni.Description = "Writing, scriptwriting, creation processes, creation tools."

		var alu Education
		alu.BeginAt = "2020"
		alu.EndAt = "2022"
		alu.SchoolName = "Alura"
		alu.Matery = "Docker, Linux, Go, UI-UX, JavaScript, Owasp Security"
		alu.Duration = "2 years"
		alu.Description = "Some Courses"

		e = []Education{
			ifsp,
			ads,
			uni,
			alu,
		}
		return
	}

	var ifsp Education
	ifsp.BeginAt = "2014"
	ifsp.EndAt = "2016"
	ifsp.SchoolName = "IFSP"
	ifsp.Matery = "Curso técnico profissionalizante integrado ao ensino médio - Manutenção e Suporte a Informática"
	ifsp.Duration = "3 anos"
	ifsp.Description = "Ensino médio e técnico integrado, UML, banco de dados, portugol, linux servers."

	var ads Education
	ads.BeginAt = "2017"
	ads.EndAt = "2017"
	ads.SchoolName = "IFSP"
	ads.Matery = "Análise e Desenvolvimento de Sistemas - Incomplete"
	ads.Duration = "1 ano"
	ads.Description = "Introdução a estatistica, reforço de matérias vistas durante o ensino técnico."

	var uni Education
	uni.BeginAt = "2018"
	uni.EndAt = "2019"
	uni.SchoolName = "Unifev"
	uni.Matery = "Publicidade e Propaganda - Icompleto"
	uni.Duration = "2 anos"
	uni.Description = "Escrita, roteirização, processos de criação, ferramentas para criação."

	var alu Education
	alu.BeginAt = "2020"
	alu.EndAt = "2022"
	alu.SchoolName = "Alura"
	alu.Matery = "Docker, Linux, Go, UI-UX, JavaScript, Owasp Security"
	alu.Duration = "Alguns cursos"
	alu.Description = "Owasp Security, Go, Linux, Docker, UI-UX, JavaScript"

	e = []Education{
		ifsp,
		ads,
		uni,
		alu,
	}

	return
}

func Experiences(l string) (e []Experience) {

	if l == "en" {

		var silk Experience
		silk.BeginAt = "jan 2016"
		silk.EndAt = "oct 2017"
		silk.Vocation = "Final Art"
		silk.BusinessName = "Silk Cromia"
		silk.Duration = "1 year"
		silk.Description = "Creation of design for different types of media."

		var flash Experience
		flash.BeginAt = "oct 2017"
		flash.EndAt = "may 2019"
		flash.Vocation = "Technical Support"
		flash.BusinessName = "Flash Net Brasil"
		flash.Duration = "1 year and 8 months"
		flash.Description = "Noc, customer support, monitoring and configuration of network equipment."

		var rosa Experience
		rosa.BeginAt = "oct 2019"
		rosa.EndAt = "feb 2020"
		rosa.Vocation = "IT Analyst"
		rosa.BusinessName = "Rosa Mística"
		rosa.Duration = "5 months"
		rosa.Description = "Analysis, management and maintenance of IT infrastructure, networks and user support, main office, " +
			"branches and group companies. (pharmacies, clinic and cremation oven industry)"

		var cofer Experience
		cofer.BeginAt = "feb 2020"
		cofer.EndAt = "apr 2021"
		cofer.Vocation = "IT Assistant"
		cofer.BusinessName = "COFERPOL"
		cofer.Duration = "1 year and 3 months"
		cofer.Description = "User support, hardware maintenance, network administration, assistance with marketing and administrative processes."

		var cofdev Experience
		cofdev.BeginAt = "may 2021"
		cofdev.EndAt = "feb 2023"
		cofdev.Vocation = "Software Enginner"
		cofdev.BusinessName = "COFERPOL"
		cofdev.Duration = "1 year and 10 months"
		cofdev.Description = "Go, DevOps culture, UML, servidores bare-metal, nginx, docker, TDD, POO, Cloudflare, Firebase e Postgres."

		var eco Experience
		eco.BeginAt = "feb 2023"
		eco.EndAt = "at the time"
		eco.Vocation = "Print on demmand"
		eco.BusinessName = "e-Commerce"
		eco.Duration = "8 months"
		eco.Description = "Knowledge about legal entities, traffic management, SEO, analytics, affiliate systems, sales funnel, and marketing."

		e = []Experience{
			eco,
			cofdev,
			cofer,
			rosa,
			flash,
			silk,
		}
		return
	}

	var silk Experience
	silk.BeginAt = "jan 2016"
	silk.EndAt = "out 2017"
	silk.Vocation = "Arte Finalista"
	silk.BusinessName = "Silk Cromia"
	silk.Duration = "1 ano"
	silk.Description = "Criação de artes para diversos tipos de mídia."

	var flash Experience
	flash.BeginAt = "out 2017"
	flash.EndAt = "mai 2019"
	flash.Vocation = "Suporte Técnico"
	flash.BusinessName = "Flash Net Brasil"
	flash.Duration = "1 ano e 8 meses"
	flash.Description = "Noc, suporte ao cliente, monitoramento e configurações de equipamentos de rede."

	var rosa Experience
	rosa.BeginAt = "out 2019"
	rosa.EndAt = "fev 2020"
	rosa.Vocation = "Analista de TI"
	rosa.BusinessName = "Rosa Mística"
	rosa.Duration = "5 meses"
	rosa.Description = "Análise, gestão e manutenção da infraestrutura de T.I, redes e suporte ao usuário, matriz, " +
		"filiais e empresas do grupo. (farmácias, clínica e indústria de fornos de cremação)"
	//TODO retirar do curri fisico

	var cofer Experience
	cofer.BeginAt = "fev 2020 "
	cofer.EndAt = "abr 2021"
	cofer.Vocation = "Assistente de TI"
	cofer.BusinessName = "COFERPOL"
	cofer.Duration = "1 ano e 3 meses"
	cofer.Description = "Suporte ao usuário, manutenção em hardware, administração de redes, auxílio em	marketing e processos administrativos."

	var cofdev Experience
	cofdev.BeginAt = "mai 2021"
	cofdev.EndAt = "fev 2023"
	cofdev.Vocation = "Desenvolvedor de Software"
	cofdev.BusinessName = "COFERPOL"
	cofdev.Duration = "1 ano e 10 meses"
	cofdev.Description = "Go, cultura DevOps, modelagem, servidores bare-metal, nginx, docker, TDD, POO, Cloudflare, Firebase e Postgres."

	var eco Experience
	eco.BeginAt = "fev 2023"
	eco.EndAt = "Até o momento"
	eco.Vocation = "Print on demmand"
	eco.BusinessName = "eCommerce"
	eco.Duration = "8 meses"
	eco.Description = "Conhecimentos sobre pessoa jurídica, gestão de tráfego, SEO, analytics, sistema de afiliados, funil de vendas e marketing."

	e = []Experience{
		eco,
		cofdev,
		cofer,
		rosa,
		flash,
		silk,
	}

	return
}

func Events(l string) (e []Event) {

	if l == "en" {

		var fli14 Event
		fli14.Name = "FLISOL (Latin American Free Software Installation Festival)."
		fli14.Type = "Congress"
		fli14.Year = 2014

		var fli15 Event
		fli15.Name = "FLISOL (Latin American Free Software Installation Festival)."
		fli15.Type = "Congress"
		fli15.Year = 2015

		var php Event
		php.Name = "PHP Light Day"
		php.Type = "Seminary"
		php.Year = 2016

		var gen Event
		gen.Name = "Genuino Day"
		gen.Type = "Congress"
		gen.Year = 2016

		var fli16 Event
		fli16.Name = "FLISOL (Latin American Free Software Installation Festival)."
		fli16.Type = "Congress"
		fli16.Year = 2016

		var con Event
		con.Name = "CONICT (Innovation, Science and Technology Congress)"
		con.Type = "Congress"
		con.Year = 2016

		e = []Event{
			con,
			fli16,
			gen,
			php,
			fli15,
			fli14,
		}
		return
	}

	var fli14 Event
	fli14.Name = "FLISOL (Festival Latino Americano de Instalação de Software Livre)."
	fli14.Type = "Congresso"
	fli14.Year = 2014

	var fli15 Event
	fli15.Name = "FLISOL (Festival Latino Americano de Instalação de Software Livre)."
	fli15.Type = "Congresso"
	fli15.Year = 2015

	var php Event
	php.Name = "PHP Light Day"
	php.Type = "Seminário"
	php.Year = 2016

	var gen Event
	gen.Name = "Genuino Day"
	gen.Type = "Congresso"
	gen.Year = 2016

	var fli16 Event
	fli16.Name = "FLISOL (Festival Latino Americano de Instalação de Software Livre)."
	fli16.Type = "Congresso"
	fli16.Year = 2016

	var con Event
	con.Name = "CONICT (Congresso de Inovação, Ciência e Tecnologia)"
	con.Type = "Congresso"
	con.Year = 2016

	e = []Event{
		con,
		fli16,
		gen,
		php,
		fli15,
		fli14,
	}

	return
}

func Articles(l string) (e []Article) {

	if l == "en" {

		var tcc Article
		tcc.Name = "Design and Development of a Humanoid Robot."
		tcc.Authors = "SILVA, G. B.; CORREIA JUNIOR, D. ; MARTINS, Osvandre. Alves"
		tcc.Type = "CONICT"
		tcc.Year = 2016

		var cot Article
		cot.Name = "Design and Development of a Humanoid Robot."
		cot.Authors = "SILVA, G. B.; CORREIA JUNIOR, D. ; MARTINS, Osvandre. Alves"
		cot.Type = "COTESI"
		cot.Year = 2016

		var con Article
		con.Name = "Design and Development of a Humanoid Robot."
		con.Authors = "SILVA, G. B.; CORREIA JUNIOR, D. ; MARTINS, Osvandre. Alves"
		con.Type = "Final paper"
		con.Year = 2016

		e = []Article{
			con,
			cot,
			tcc,
		}
		return
	}

	var tcc Article
	tcc.Name = "Concepção e Desenvolvimento de um Robô Humanoide."
	tcc.Authors = "SILVA, G. B.; CORREIA JUNIOR, D. ; MARTINS, Osvandre. Alves"
	tcc.Type = "CONICT"
	tcc.Year = 2016

	var cot Article
	cot.Name = "Concepção e Desenvolvimento de um Robô Humanoide."
	cot.Authors = "SILVA, G. B.; CORREIA JUNIOR, D. ; MARTINS, Osvandre. Alves"
	cot.Type = "COTESI"
	cot.Year = 2016

	var con Article
	con.Name = "Concepção e Desenvolvimento de um Robô Humanoide."
	con.Authors = "SILVA, G. B.; CORREIA JUNIOR, D. ; MARTINS, Osvandre. Alves"
	con.Type = "Trabalho de Conclusão de Curso"
	con.Year = 2016

	e = []Article{
		con,
		cot,
		tcc,
	}

	return
}
