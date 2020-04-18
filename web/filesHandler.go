package web

import (
	"ChangeInspector/gitlog"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

/*FilesHandler ...*/
type FilesHandler struct {
	gitLog *gitlog.GitLog
}

func (handler FilesHandler) register(router *mux.Router) {

	router.HandleFunc("/files/{fileName}/commits", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fileName := strings.ReplaceAll(vars["fileName"], "$", "/")
		commits := handler.gitLog.GetFileCommits(fileName)
		json.NewEncoder(w).Encode(commits)
	})

	router.HandleFunc("/files/{fileName}/detail", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fileName := strings.ReplaceAll(vars["fileName"], "$", "/")
		fileInfo := handler.gitLog.GetFileInfo(fileName)
		json.NewEncoder(w).Encode(fileInfo)
	})
}
