package app

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/mini-eggs/home-server/app/database"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type routeType struct {
	path   string
	handle func(http.ResponseWriter, *http.Request)
}

type appType struct {
	path   string
	routes []routeType
}

var apps = [3]appType{
	{path: "", routes: []routeType{
		{path: "/", handle: defaultRoute("Pesto's House")},
	}},
	{path: "/snack-chan", routes: []routeType{
		{path: "", handle: defaultRoute("Snack Chan")},
	}},
	{path: "/testing", routes: []routeType{
		{path: "/database", handle: testDatabaseConnect},
	}},
}

func testDatabaseConnect(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	defer db.Close()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Database is not working as expected."))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Database is working as expected."))
}

func defaultRoute(title string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		type dataType struct {
			Title string
		}
		data := dataType{
			Title: title,
		}

		errorOut := func(err error) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Could not locate page."))
		}

		template, parseTemplateError := template.ParseFiles("templates/default/index.html")
		if parseTemplateError != nil {
			errorOut(parseTemplateError)
			return
		}

		var templateBuffer bytes.Buffer
		toStringErr := template.Execute(&templateBuffer, data)
		if toStringErr != nil {
			errorOut(toStringErr)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(templateBuffer.String()))
	}
}

func registerRoutes(router *mux.Router) {
	for i := 0; i < len(apps); i++ {
		current := apps[i]
		aRouter := router.PathPrefix(current.path).Subrouter()
		for e := 0; e < len(current.routes); e++ {
			nestedRoute := current.routes[e]
			aRouter.HandleFunc(nestedRoute.path, nestedRoute.handle)
		}
	}
}
