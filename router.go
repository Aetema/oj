package main

import (
	"net/http"

	"github.com/Miloas/oj/controller"
)

//Routes : router
func Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/problem/add", controller.HandleAddProblem)
	mux.HandleFunc("/problem/remove", controller.HandleRemoveProblem)
	mux.HandleFunc("/problem/update", controller.HandleUpdateProblem)
	mux.HandleFunc("/problem/submit", controller.HandleSubmitCode)
	mux.HandleFunc("/problem", controller.HandleProblem)
	mux.HandleFunc("/status", controller.HandleStatus)
	mux.HandleFunc("/signup", controller.HandleSignup)
	mux.HandleFunc("/signout", controller.HandleSignout)
	mux.HandleFunc("/signin", controller.HandleSignin)
	//display problems list
	mux.HandleFunc("/", controller.HandleHome)
	//add static file server for include static files
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	return mux
}
