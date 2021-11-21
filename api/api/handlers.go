package api

import (
	"fmt"
	"github.com/tonnytg/ccenter/pkg/gcp/project"
	"html/template"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "http://localhost:8080/project\n")
	fmt.Fprintf(w, "http://localhost:8080/projecthtml")
}

type NewAggPage struct {
	Title string
	News  string
}

func getProjects(w http.ResponseWriter, r *http.Request) {
	log.Printf("rootHandler accessed: %s", r.URL.Path)

	p, err := project.Get()
	if err != nil {
		log.Printf("Error getting projects: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, project := range p.Projects {
		fmt.Fprintf(w, "ProjectID[%s]: %s\n", project.ProjectNumber, project.Name)
	}
}

func getProjectsHTML(w http.ResponseWriter, r *http.Request) {

	pj, err := project.Get()
	if err != nil {
		log.Printf("Error getting projects: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//p := NewAggPage{Title: "Hi", News: "ok news"}
	t, err := template.ParseFiles("./template/basic.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, pj)
	fmt.Println(err)
}
