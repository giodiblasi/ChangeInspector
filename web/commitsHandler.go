package web

import (
	"ChangeInspector/gitlog"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

/*CommitsHandler ...*/
type CommitsHandler struct {
	gitLog gitlog.GitLog
}

func (handler CommitsHandler) register(router *mux.Router) {
	router.HandleFunc("/commits/{hash}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		hash := vars["hash"]
		commit := handler.gitLog.Commits[hash]
		json.NewEncoder(w).Encode(commit)
	})
}
