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
	if !strings.HasPrefix(u.Path, "/problem/add") || !strings.HasPrefix(u.Path, "/problem/remove") || !strings.HasPrefix(u.Path, "/problem/update") || !strings.HasPrefix(u.Path, "/problem/submit") {
		next(w, r)
	} else {
		if controller.GetIsadmin(r) {
			next(w, r)
		} else {
			http.Error(w, "Not Authorized", 401)
		}
	}
}
