package web

import (
	"ChangeInspector/logservice"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

/*FilesHandler ...*/
type FilesHandler struct {
	logService *logservice.LogService
}

func (handler FilesHandler) register(router *mux.Router) {

	router.HandleFunc("/files/{fileName}/commits", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fileName := strings.ReplaceAll(vars["fileName"], "$", "/")
		commits := handler.logService.GitLog.GetFileCommits(fileName)
		json.NewEncoder(w).Encode(commits)
	})

	router.HandleFunc("/files/{fileName}/detail", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fileName := strings.ReplaceAll(vars["fileName"], "$", "/")
		fileInfo := handler.logService.GitLog.GetFileInfo(fileName)
		json.NewEncoder(w).Encode(fileInfo)
	})
}
