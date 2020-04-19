package web

import (
	"ChangeInspector/logservice"
	"fmt"
	"net/http"
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
		} else {
			handler.logService.Update(before, after)
		}
	}).Methods("PUT")
}
