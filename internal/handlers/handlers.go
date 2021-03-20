package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jfk23/gobookings/internal/config"
	"github.com/jfk23/gobookings/internal/model"
	"github.com/jfk23/gobookings/internal/render"

)

var Repo *Repository

type Repository struct {
	ConfigSetting *config.AppConfig
}

func CreateNewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		ConfigSetting: a,
	}
}

func SetHandler(r *Repository) {
	Repo = r
}

//Home is page for /
func (re *Repository) Home(rw http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(rw, r, "home.page.html", &model.TemplateData{})

	remoteIP := r.RemoteAddr

	re.ConfigSetting.Session.Put(r.Context(), "remote_ip", remoteIP)

}

func (re *Repository) Reservation(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, r, "search.page.html", &model.TemplateData{})
}

func (re *Repository) Generals(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, r, "generals.page.html", &model.TemplateData{})
}

func (re *Repository) Majors(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, r, "majors.page.html", &model.TemplateData{})
}

func (re *Repository) Contact(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, r, "contact.page.html", &model.TemplateData{})
}

func (re *Repository) Search(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, r, "search.page.html", &model.TemplateData{
	})
}

func (re *Repository) PostSearch(rw http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	rw.Write([]byte(fmt.Sprintf("start dat is %s, and end date is %s", start, end)))

}

type responseJson struct {
	OK bool `json:"ok"`
	Message string `json:"message"`

}

func (re *Repository) PostSearchJson(rw http.ResponseWriter, r *http.Request) {
	var resp = responseJson {
		OK : true,
		Message: "hi there",
	}

	out, err := json.MarshalIndent(resp, " ", "    ")

	if err !=nil {
		fmt.Println(err)
	}

	fmt.Println(string(out))

	rw.Header().Add("Content-Type", "application/json")
	rw.Write(out)
	
}


// About is page for /
func (re *Repository) About(rw http.ResponseWriter, r *http.Request) {
	remoteIP := re.ConfigSetting.Session.GetString(r.Context(), "remote_ip")

	var stringData = map[string]string{"test": "Hello Good day!", "remote_ip": remoteIP}
	render.RenderTemplate(rw, r, "about.page.html", &model.TemplateData{
		StringMap: stringData,
	})

}
