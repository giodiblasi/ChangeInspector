package web

import (
	"ChangeInspector/commits"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

/*DetailHandler ...*/
type DetailHandler struct {
	filesInfo *commits.FileInfos
}

func (handler DetailHandler) register(router *mux.Router) {

	router.HandleFunc("/detail/{fileName}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fileName := strings.ReplaceAll(vars["fileName"], "$", "/")
		fileInfo := (*handler.filesInfo)[fileName]
		json.NewEncoder(w).Encode(fileInfo)
	})
}
