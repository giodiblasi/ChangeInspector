package web

import (
	"ChangeInspector/commits"
	"ChangeInspector/sortedcommits"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gorilla/mux"
)

/*SortHandler ...*/
type SortHandler struct {
	filesInfo *commits.FileInfos
}

func getResult(query url.Values, result sortedcommits.GoogleChartBarResult) sortedcommits.GoogleChartBarResult {
	take, ok := query["take"]
	if ok {
		i, _ := strconv.Atoi(take[0])
		return result[:i]
	}
	return result
}

func (handler SortHandler) register(router *mux.Router) {
	arrayInfo := make([]sortedcommits.OrderableFileInfo, 0)
	for fileName, fileInfo := range *handler.filesInfo {
		arrayInfo = append(arrayInfo, sortedcommits.OrderableFileInfo{
			FileName: fileName,
			Info:     fileInfo,
		})
	}

	router.HandleFunc("/sort/commits", func(w http.ResponseWriter, r *http.Request) {
		result := sortedcommits.Sort(arrayInfo, sortedcommits.ByCommits{})
		query := r.URL.Query()
		json.NewEncoder(w).Encode(getResult(query, result))
	})

	router.HandleFunc("/sort/changes", func(w http.ResponseWriter, r *http.Request) {
		result := sortedcommits.Sort(arrayInfo, sortedcommits.ByChanges{})
		query := r.URL.Query()
		json.NewEncoder(w).Encode(getResult(query, result))
	})

}
