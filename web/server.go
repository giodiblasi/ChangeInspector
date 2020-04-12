package web

import (
	"ChangeInspector/commits"
	"encoding/json"
	"html/template"
	"net/http"
)

var tpl = template.Must(template.ParseFiles("web/index.html"))

type pageModel struct {
	Title     string
	FilesInfo string
}

func indexHandler(model pageModel) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, model)
	}
}

/*StartServer ...*/
func StartServer(filesInfo commits.FileInfos) {
	data, _ := json.Marshal(filesInfo)
	model := pageModel{Title: "ChangeInspector", FilesInfo: string(data)}

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler(model))
	SortHandler{filesInfo: filesInfo}.register(mux)

	staticFileServer := http.FileServer(http.Dir("web/assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", staticFileServer))

	http.ListenAndServe(":3000", mux)
}
