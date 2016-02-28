package controller

import (
	"net/http"

	"github.com/Miloas/oj/model"

	"gopkg.in/mgo.v2/bson"
)

type contestResult struct {
	User      model.User
	Problems  []model.Problem
	ContestID string
	Islogin   bool
}

//HandleContest : handle contest page ? id = :id
func HandleContest(w http.ResponseWriter, r *http.Request) {
	//cid means contest cid , problem id will be pid there
	cid := r.URL.Query().Get("cid")
	session := getMongoS()
	defer session.Close()
	c := session.DB("oj").C("contests")
	userCol := session.DB("oj").C("user")
	user := model.User{}
	contests := []model.Contest{}
	userCol.Find(bson.M{"username": GetLoginUser(r).Username}).One(&user)
	c.Find(bson.M{"contestid": cid}).All(&contests)
	if len(contests) > 0 {
		problemIDs := contests[0].ContestProblems
		problemCol := session.DB("oj").C("problems")
		problems := []model.Problem{}
		problemsRet := []model.Problem{}
		for _, v := range problemIDs {
			problemCol.Find(bson.M{"id": v}).All(&problemsRet)
			if len(problemsRet) > 0 {
				problems = append(problems, problemsRet[0])
			}
		}
		Render.HTML(w, http.StatusOK, "contest", contestResult{user, problems, cid, GetIslogin(r)})
	} else {
		http.Redirect(w, r, "/contests", http.StatusNotFound)
	}
}
