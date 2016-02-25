package controller

import (
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/Miloas/oj/model"
)

type boardPageStruct struct {
	Problems  []int
	Users     []model.User
	ContestID string
	Islogin   bool
}

//HandleContestBoard : handle contest board page
func HandleContestBoard(w http.ResponseWriter, r *http.Request) {
	cid := r.URL.Query().Get("cid")
	session := getMongoS()
	defer session.Close()
	c := session.DB("oj").C("user")
	contestCol := session.DB("oj").C("contests")
	contest := model.Contest{}
	contestCol.Find(bson.M{"contestid": cid}).One(&contest)
	problems := []int{}
	for i := range contest.ContestProblems {
		problems = append(problems, i)
	}
	result := []model.User{}
	c.Find(bson.M{"joinedcontest": cid}).Sort("-contesttotalaced", "contesttotaltime").All(&result)
	Render.HTML(w, http.StatusFound, "contestBoard", boardPageStruct{problems, result, cid, GetIslogin(r)})
}
