package main

import (
	"net/http"

	"github.com/Miloas/oj/controller"
)

//Routes : router
func Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/problem/add", controller.HandleAddProblem)       // need admin
	mux.HandleFunc("/problem/remove", controller.HandleRemoveProblem) // need admin
	mux.HandleFunc("/problem/update", controller.HandleUpdateProblem) // need admin
	mux.HandleFunc("/problem/submit", controller.HandleSubmitCode)
	mux.HandleFunc("/problem", controller.HandleProblem)
	mux.HandleFunc("/status", controller.HandleStatus)
	mux.HandleFunc("/signup", controller.HandleSignup)
	mux.HandleFunc("/signout", controller.HandleSignout)
	mux.HandleFunc("/signin", controller.HandleSignin)
	mux.HandleFunc("/contests", controller.HandleContests)
	mux.HandleFunc("/contest/add", controller.HandleAddContest)                // need admin
	mux.HandleFunc("/contest/remove", controller.HandleRemoveContest)          // need admin
	mux.HandleFunc("/contest/update", controller.HandleUpdateContest)          // need admin
	mux.HandleFunc("/contest", controller.HandleContest)                       //during time or admin
	mux.HandleFunc("/contest/problems/submit", controller.HandleContestSubmit) //during time or admin
	mux.HandleFunc("/contest/problems", controller.HandleContestProblems)      //during time or admin
	mux.HandleFunc("/contest/status", controller.HandleContestStatus)          //during time or admin
	mux.HandleFunc("/contest/board", controller.HandleContestBoard)            //during time or admin
	mux.HandleFunc("/error", controller.HandleError)
	//display problems list
	mux.HandleFunc("/", controller.HandleHome)
	//add static file server for include static files
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	return mux
}
