package web

import (
	"ChangeInspector/logservice"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var tpl = template.Must(template.ParseFiles("web/index.html"))

type pageModel struct {
	Title     string
	StartDate string
	EndDate   string
}

func indexHandler(logService *logservice.LogService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		model := pageModel{
			Title:     "Changes Inspector",
			StartDate: logService.GitLog.After.Format("2006-01-02"),
			EndDate:   logService.GitLog.Before.Format("2006-01-02"),
		}
		tpl.Execute(w, model)
	}
}

/*StartServer ...*/
func StartServer(logService *logservice.LogService) {
	router := mux.NewRouter()

	router.HandleFunc("/", indexHandler(logService))
	SortHandler{logService}.register(router)
	FilesHandler{logService}.register(router)
	CommitsHandler{logService}.register(router)
	LogHandler{logService}.register(router)

	staticFileServer := http.FileServer(http.Dir("web/assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", staticFileServer))
	fmt.Println("server starts to http://localhost:3000")
	http.ListenAndServe(":3000", router)
}
