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

func getProjects(w http.ResponseWriter, r *http.Request) {
	log.Printf("rootHandler accessed: %s", r.URL.Path)

	p, err := project.GetID()
	if err != nil {
        log.Printf("Error getting projects: %v", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

	for i, project := range p {
        fmt.Fprintf(w, "ProjectID[%d]: %v\n", i, project)
    }
}

func getProjectsHTML(w http.ResponseWriter, r *http.Request) {

	ListProjects := template.Must(template.New("ListProjects").Parse(`
	
    <html>
		<head>
			<title>List Projects</title>
		</head>
		<body>
			<h1>List Projects</h1>
			<table border="1">
				<tr><th>ProjectID</th></tr>
				{{range .}}
					<tr><td>{{.}}</td></tr>
				{{end}}
			</table>
		</body>
    </html>
    `))


	log.Printf("rootHandler accessed: %s", r.URL.Path)

	p, err := project.GetID()
	if err != nil {
		log.Printf("Error getting projects: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//for i, project := range p {
	//	fmt.Fprintf(w, "ProjectID[%d]: %v\n", i, project)
	//}

	if err := ListProjects.Execute(w, p); err != nil {
        log.Printf("Error executing template: %v", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
}