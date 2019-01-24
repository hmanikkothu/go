package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	indexTmpl = parseTemplate("index.html")
	page1Tmpl = parseTemplate("page1.html")
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9999"
	}
	registerHandlers()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func registerHandlers() {
	// Use gorilla/mux for rich routing.
	// See http://www.gorillatoolkit.org/pkg/mux
	r := mux.NewRouter()

	r.Handle("/", http.RedirectHandler("/index", http.StatusFound))
	r.Methods("GET").Path("/index").
		Handler(appHandler(indexHandler))
	r.Methods("GET").Path("/page1").
		Handler(appHandler(page1Handler))

	// Respond to App Engine and Compute Engine health checks.
	// Indicate the server is healthy.
	r.Methods("GET").Path("/_ah/health").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ok"))
		})

	// [START request_logging]
	// Delegate all of the HTTP routing and serving to the gorilla/mux router.
	// Log all requests using the standard Apache format.
	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, r))
	// [END request_logging]

}

// index page.
func indexHandler(w http.ResponseWriter, r *http.Request) *appError {
	data := struct {
		Title       string
		Description string
	}{
		"Index Page",
		"content - runtime",
	}

	return indexTmpl.Execute(w, r, data)
}

// page1 page.
func page1Handler(w http.ResponseWriter, r *http.Request) *appError {
	data := struct {
		Title       string
		Description string
	}{
		"Page1 Page",
		"content - runtime",
	}

	return page1Tmpl.Execute(w, r, data)
}
