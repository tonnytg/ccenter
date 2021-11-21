package api

import (
	"net/http"
)

func Start() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/project", getProjects)
	http.HandleFunc("/projecthtml", getProjectsHTML)
	http.ListenAndServe(":8080", nil)
}
