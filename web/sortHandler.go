package web

import (
	"ChangeInspector/commits"
	"ChangeInspector/sortedcommits"
	"encoding/json"
	"net/http"
)

/*SortHandler ...*/
type SortHandler struct {
	filesInfo commits.FileInfos
}

func (handler SortHandler) register(mux *http.ServeMux) {
	arrayInfo := make([]sortedcommits.OrderableFileInfo, 0)
	for fileName, fileInfo := range handler.filesInfo {
		arrayInfo = append(arrayInfo, sortedcommits.OrderableFileInfo{
			FileName: fileName,
			Info:     fileInfo,
		})
	}

	mux.HandleFunc("/sort/commits", func(w http.ResponseWriter, r *http.Request) {
		result := sortedcommits.Sort(arrayInfo, sortedcommits.ByCommits{})
		json.NewEncoder(w).Encode(result)
	})

	mux.HandleFunc("/sort/changes", func(w http.ResponseWriter, r *http.Request) {
		result := sortedcommits.Sort(arrayInfo, sortedcommits.ByChanges{})
		json.NewEncoder(w).Encode(result)
	})

}
