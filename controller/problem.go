package controller

import (
	"net/http"
	"strconv"

	"github.com/Miloas/oj/model"

	"gopkg.in/mgo.v2/bson"
)

type problemResult struct {
	Problem model.Problem
	Islogin bool
}

//HandleProblem : handle "/problem?id=:id"
func HandleProblem(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	_, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	session := getMongoS()
	defer session.Close()
	c := session.DB("oj").C("problems")
	var ret model.Problem
	err = c.Find(bson.M{"id": id}).One(&ret)
	if err != nil {
		panic(err)
	}
	Render.HTML(w, http.StatusOK, "problem", problemResult{ret, GetIslogin(r)})
}
