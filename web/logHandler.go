package web

import (
	"ChangeInspector/logservice"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

/*LogHandler ...*/
type LogHandler struct {
	logService *logservice.LogService
}

func (handler LogHandler) register(router *mux.Router) {
	router.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		layout := "2006-01-02"
		var err error
		var after, before time.Time
		before, err = time.Parse(layout, query["before"][0])
		after, err = time.Parse(layout, query["after"][0])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, err.Error())
			return
		}
		if before.Sub(after).Milliseconds() < 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "end date have to be greater that start date")
			return
		}
		handler.logService.Update(before, after)

	}).Methods("PUT")

	router.HandleFunc("/filter/{file}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		file := strings.ReplaceAll(vars["file"], "$", "/")

		handler.logService.AddFileToFilter(file)

	}).Methods("POST")

	router.HandleFunc("/filter/{file}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		file := strings.ReplaceAll(vars["file"], "$", "/")

		ok := handler.logService.RemoveFileFromFilter(file)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

	}).Methods("DELETE")

	router.HandleFunc("/filter", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(handler.logService.Filter)

	}).Methods("GET")
}
