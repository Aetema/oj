package controller

import "net/http"

//HandleContest : handle contest page ? id = :id
func HandleContest(w http.ResponseWriter, r *http.Request) {
	Render.HTML(w, http.StatusOK, "contest", nil)
}
