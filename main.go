package main

import (
	"github.com/Miloas/oj/middleware"

	"github.com/codegangsta/negroni"
)

func main() {
	mux := Routes()
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.HandlerFunc(middleware.Permission),
		negroni.NewLogger(),
	)
	n.UseHandler(mux)
	n.Run(":8080")
}
