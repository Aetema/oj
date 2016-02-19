package controller

import (
	"net/http"

	"gopkg.in/boj/redistore.v1"
)

//HandleSignout : handle signout action
func HandleSignout(w http.ResponseWriter, r *http.Request) {
	store, err := redistore.NewRediStore(10, "tcp", ":6379", "", []byte("secret-key"))
	if err != nil {
		panic(err)
	}
	defer store.Close()
	// Get a session.
	accountSession, _ := store.Get(r, "user")
	accountSession.Options.MaxAge = -1
	// Save.
	accountSession.Save(r, w)
	http.Redirect(w, r, "/", http.StatusFound)
}
