package web

import (
	"ChangeInspector/commits"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var tpl = template.Must(template.ParseFiles("web/index.html"))

type pageModel struct {
	Title string
}

func indexHandler(model pageModel) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, model)
	}
}

/*StartServer ...*/
func StartServer(filesInfo commits.FileInfos) {
	model := pageModel{Title: "ChangeInspector"}

	router := mux.NewRouter()

	router.HandleFunc("/", indexHandler(model))

	SortHandler{filesInfo: &filesInfo}.register(router)
	DetailHandler{filesInfo: &filesInfo}.register(router)

	staticFileServer := http.FileServer(http.Dir("web/assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", staticFileServer))

	http.ListenAndServe(":3000", router)
}
