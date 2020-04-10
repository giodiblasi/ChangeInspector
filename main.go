package main

import (
	"ChangeInspector/commits"
	"ChangeInspector/sortedcommits"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"text/template"
)

func getCommitsResult(path string) string {
	cmd := exec.Command("bash", "stat", path)
	out, _ := cmd.CombinedOutput()
	return string(out)
}

var tpl = template.Must(template.ParseFiles("index.html"))

type pageModel struct {
	Title     string
	FilesInfo string
}

func getHandler(fileInfo commits.FileInfos) func(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(fileInfo)
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, pageModel{Title: "Yeah", FilesInfo: string(data)})
	}
}

func main() {
	var gitFolder string = os.Args[1]
	var result string = getCommitsResult(gitFolder)
	var fileInfos commits.FileInfos = commits.Parse(result)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", getHandler(fileInfos))
	http.ListenAndServe(":"+port, mux)

	for _, fileInfo := range sortedcommits.GetSorted(fileInfos, sortedcommits.ByCommits{}) {
		fmt.Println("file:", fileInfo.FileName, "changed", fileInfo.Info.TotalChanges, "times in", len(fileInfo.Info.Commits), "commits")
	}

}
