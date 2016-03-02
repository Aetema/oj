package controller

import (
	"net/http"

	"github.com/Miloas/oj/model"

	"gopkg.in/boj/redistore.v1"
	"gopkg.in/mgo.v2/bson"
)

//HandleSignin : handle signin page
func HandleSignin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		session := getMongoS()
		c := session.DB("oj").C("user")
		result := []model.User{}
		c.Find(bson.M{"username": r.Form["loginUsername"][0]}).All(&result)
		store, err := redistore.NewRediStore(10, "tcp", ":6379", "", []byte("secret-key"))
		if err != nil {
			panic(err)
		}
		defer store.Close()
		if len(result) <= 0 || cryptoPassword(r.Form["loginPassword"][0]) != result[0].Password {
			accountSession, _ := store.Get(r, "info")
			accountSession.Values["loginInfo"] = "账号密码错误."
			accountSession.Save(r, w)
		} else {
			accountSession, _ := store.Get(r, "user")
			accountSession.Values["currentuser"] = &result[0]
			accountSession.Save(r, w)

			accountInfoSession, _ := store.Get(r, "info")
			accountInfoSession.Values["loginInfo"] = "login successful."
			accountInfoSession.Save(r, w)
		}
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
