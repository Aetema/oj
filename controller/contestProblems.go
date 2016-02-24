package controller

import (
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/Miloas/oj/model"
)

type contestProblemResult struct {
	Problem          model.Problem
	ContestID        string
	ContestProblemID string
	Islogin          bool
}

//HandleContestProblems : handle contest list problems page, (/contest/problems?cid=:cid&&pid=:pid)
func HandleContestProblems(w http.ResponseWriter, r *http.Request) {
	cid := r.URL.Query().Get("cid")
	pid := r.URL.Query().Get("pid")
	session := getMongoS()
	defer session.Close()
	contestCol := session.DB("oj").C("contests")
	problemCol := session.DB("oj").C("problems")
	contest := []model.Contest{}
	problem := []model.Problem{}
	contestCol.Find(bson.M{"contestid": cid}).All(&contest)
	if len(contest) > 0 {
		a2i := map[string]int{"A": 0, "B": 1, "C": 2, "D": 3, "E": 4, "F": 5, "G": 6, "H": 7, "I": 8, "J": 9, "K": 10, "L": 11, "M": 12, "N": 13, "O": 14, "P": 15, "Q": 16, "R": 17, "S": 18, "T": 19, "U": 20, "V": 21, "W": 22, "X": 23, "Y": 24, "Z": 25}
		if v, ok := a2i[pid]; ok {
			problemCol.Find(bson.M{"id": contest[0].ContestProblems[v]}).All(&problem)
			if len(problem) > 0 {
				Render.HTML(w, http.StatusFound, "contestProblem", contestProblemResult{problem[0], cid, pid, GetIslogin(r)})
			} else {
				http.Redirect(w, r, "/error", 401)
			}
		} else {
			http.Redirect(w, r, "/error", 401)
		}
	} else {
		http.Redirect(w, r, "/error", 401)
	}
}
