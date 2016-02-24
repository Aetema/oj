package middleware

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/Miloas/oj/controller"
)

//Permission :Use to check permission when user touch crud (Yet finish)
func Permission(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	u, _ := url.ParseRequestURI(r.URL.String())
	//restrict crud problems and contest (only admin)
	if !strings.HasPrefix(u.Path, "/problem/add") && !strings.HasPrefix(u.Path, "/problem/remove") && !strings.HasPrefix(u.Path, "/problem/update") &&
		!strings.HasPrefix(u.Path, "/contest/add") && !strings.HasPrefix(u.Path, "/contest/remove") && !strings.HasPrefix(u.Path, "/contest/update") {
		if strings.HasPrefix(u.Path, "/problem") || strings.HasPrefix(u.Path, "/problem/submit") {
			if controller.CheckAuth2Problem(r) {
				next(w, r)
				return
			}
			http.Redirect(w, r, "/error", 401)
			return
			// http.Error(w, "Not Authorized", 401)
		}
		if strings.HasPrefix(u.Path, "/contest") && !strings.HasPrefix(u.Path, "/contests") {
			if controller.CheckAuth2Contest(r) {
				next(w, r)
				return
			}
			http.Redirect(w, r, "/error", 401)
			return
		}
		next(w, r)
	} else {
		if controller.GetIsadmin(r) {
			next(w, r)
		}
		http.Redirect(w, r, "/error", 401)
		return
		// http.Error(w, "Not Authorized", 401)
	}

}
