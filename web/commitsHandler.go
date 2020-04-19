package web

import (
	"ChangeInspector/logservice"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

/*CommitsHandler ...*/
type CommitsHandler struct {
	logService *logservice.LogService
}

func (handler CommitsHandler) register(router *mux.Router) {
	router.HandleFunc("/commits/{hash}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		hash := vars["hash"]
		commit := handler.logService.GitLog.Commits[hash]
		json.NewEncoder(w).Encode(commit)
	})
}
