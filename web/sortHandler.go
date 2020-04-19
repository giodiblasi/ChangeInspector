package web

import (
	"ChangeInspector/logservice"
	"ChangeInspector/sorter"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gorilla/mux"
)

/*SortHandler ...*/
type SortHandler struct {
	logService *logservice.LogService
}

func getResult(query url.Values, result sorter.GoogleChartBarResult) sorter.GoogleChartBarResult {
	takeParam, ok := query["take"]
	if ok {
		take, _ := strconv.Atoi(takeParam[0])
		length := len(result)
		if length < take {
			take = length
		}
		return result[:take]
	}
	return result
}

func (handler SortHandler) register(router *mux.Router) {
	router.HandleFunc("/sort/commits", func(w http.ResponseWriter, r *http.Request) {
		result := handler.logService.SortableLogs.SortBy(sorter.ByCommits{})
		query := r.URL.Query()
		json.NewEncoder(w).Encode(getResult(query, result))
	})

	router.HandleFunc("/sort/changes", func(w http.ResponseWriter, r *http.Request) {
		result := handler.logService.SortableLogs.SortBy(sorter.ByChanges{})
		query := r.URL.Query()
		json.NewEncoder(w).Encode(getResult(query, result))
	})

}
