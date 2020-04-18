package web

import (
	"ChangeInspector/gitlog"
	"net/http"

	"github.com/gorilla/mux"
)

/*LogHandler ...*/
type LogHandler struct {
	gitLog *gitlog.GitLog
}

func (handler LogHandler) register(router *mux.Router) {
	router.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		before := query["before"][0]
		after := query["after"][0]
		handler.gitLog.Update(before, after)
	})
}
