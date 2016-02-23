package controller

import "net/http"

//HandleError : if page deny,redirect this page
func HandleError(w http.ResponseWriter, r *http.Request) {
	Render.HTML(w, 401, "error", nil)
}
