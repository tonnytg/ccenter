package project

import (
	"encoding/json"
	"fmt"
	"github.com/tonnytg/ccenter/pkg/scrap"
	"time"
)

type AllProjects struct {
	Projects []struct {
		ProjectNumber  string    `json:"projectNumber"`
		ProjectID      string    `json:"projectId"`
		LifecycleState string    `json:"lifecycleState"`
		Name           string    `json:"name"`
		CreateTime     time.Time `json:"createTime"`
	} `json:"projects"`
}

func GetID() ([]string, error) {

	// GET ServiceAccount of one project
	url := "https://cloudresourcemanager.googleapis.com/v1/projects"

	data, err := scrap.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(string(data))

	var allProjects AllProjects
	var projects []string
	json.Unmarshal(data, &allProjects)
	for i, v := range allProjects.Projects {
		fmt.Printf("Project[%d]: %s\n", i, v.ProjectID)
		projects = append(projects, v.ProjectID)
	}

	return projects, nil
}