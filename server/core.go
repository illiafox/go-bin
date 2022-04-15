package server

import (
	"gobin/server/methods"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gobin/database"
	"gobin/utils/config"
)

func Serve(db *database.Database, conf config.Host) *http.Server {
	root := mux.NewRouter()

	// static
	static := http.FileServer(http.Dir("../../shared/static"))
	shared := http.FileServer(http.Dir("../../shared"))

	root.Handle("/", static)
	root.PathPrefix("/css/").Handler(shared)
	root.PathPrefix("/images/").Handler(shared)
	root.Handle("/favicon.ico", shared)
	//
	m := methods.New(db)

	root.HandleFunc("/new", m.New)
	root.HandleFunc("/{key}", m.Get)

	//

	return &http.Server{
		Addr: "0.0.0.0:" + conf.Port,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      root, // Pass our instance of gorilla/mux in
	}
}
