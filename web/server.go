package web

import (
	"ChangeInspector/logservice"
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
func StartServer(logService *logservice.LogService) {
	model := pageModel{Title: "Change Inspector"}

	router := mux.NewRouter()

	router.HandleFunc("/", indexHandler(model))
	SortHandler{logService}.register(router)
	FilesHandler{logService}.register(router)
	CommitsHandler{logService}.register(router)
	LogHandler{logService}.register(router)

	staticFileServer := http.FileServer(http.Dir("web/assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", staticFileServer))

	http.ListenAndServe(":3001", router)
}
