package web

import (
	"ChangeInspector/gitlog"
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
func StartServer(gitLog gitlog.GitLog) {
	model := pageModel{Title: "ChangeInspector"}

	router := mux.NewRouter()

	router.HandleFunc("/", indexHandler(model))

	SortHandler{gitLog: gitLog}.register(router)
	FilesHandler{gitLog: gitLog}.register(router)
	CommitsHandler{gitLog: gitLog}.register(router)

	staticFileServer := http.FileServer(http.Dir("web/assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", staticFileServer))

	http.ListenAndServe(":3000", router)
}
