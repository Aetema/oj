package controller

import (
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/Miloas/oj/model"
	"gopkg.in/boj/redistore.v1"
)

//HandleSignup : handle signup page post info
func HandleSignup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		//	fmt.Println(r.Form["createUsername"][0], r.Form["createPassword"][0], r.Form["repeatePassword"][0])
		if r.Form["createPassword"][0] != r.Form["repeatePassword"][0] {
			//io.WriteString(w, "注册失败，两次密码不一致")

		} else {
			session := getMongoS()
			defer session.Close()
			c := session.DB("oj").C("user")
			result := []model.User{}
			c.Find(bson.M{"username": r.Form["createUsername"][0]}).All(&result)
			// fmt.Println(len(result))
			if len(result) <= 0 {
				// fmt.Println("ok")

				if len(r.Form["createUsername"][0]) >= 6 && len(r.Form["createUsername"][0]) <= 12 && len(r.Form["createPassword"][0]) >= 7 && len(r.Form["createPassword"][0]) <= 12 {
					cryptoedPassword := cryptoPassword(r.Form["createPassword"][0])
					// fmt.Println(cryptoedPassword)
					user := model.User{r.Form["createUsername"][0], cryptoedPassword, "normal", []string{}, ""}
					c.Insert(&user)
					store, err := redistore.NewRediStore(10, "tcp", ":6379", "", []byte("secret-key"))
					if err != nil {
						panic(err)
					}
					defer store.Close()
					accountSession, _ := store.Get(r, "user")
					accountSession.Values["currentuser"] = &user
					accountSession.Save(r, w)
				}
			}
		}
	}
}
