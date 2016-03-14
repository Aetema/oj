package controller

import (
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/Miloas/oj/model"
)

type boardPageStruct struct {
	Problems  []int
	Users     []model.User
	ContestID string
	Islogin   bool
}

type closedBoardPageStruct struct {
	Problems  []int
	Users     []model.ClosedUser
	ContestID string
	Islogin   bool
}

//HandleContestBoard : handle contest board page
func HandleContestBoard(w http.ResponseWriter, r *http.Request) {
	cid := r.URL.Query().Get("cid")
	session := getMongoS()
	defer session.Close()
	contestCol := session.DB("oj").C("contests")
	contest := model.Contest{}
	contestCol.Find(bson.M{"contestid": cid}).One(&contest)
	problems := []int{}
	for i := range contest.ContestProblems {
		problems = append(problems, i)
	}

	contestStartTime, _ := time.Parse("2006-01-02 15:04:05", contest.StartTime)
	contestHowlong := contest.HowLong
	contestEndTime := contestStartTime
	contestCloseTime := contestStartTime
	for i := 0; i < contestHowlong-1; i++ {
		contestCloseTime = contestCloseTime.Add(time.Hour)
	}
	for i := 0; i < contestHowlong; i++ {
		contestEndTime = contestEndTime.Add(time.Hour)
	}
	currentTime, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
	if currentTime.After(contestCloseTime) && currentTime.Before(contestEndTime) && !GetIsadmin(r) {
		result := []model.ClosedUser{}
		c := session.DB("oj").C("closeduser")
		c.Find(bson.M{"joinedcontest": cid}).Sort("-contesttotalaced", "contesttotaltime").All(&result)
		Render.HTML(w, http.StatusFound, "contestBoard", closedBoardPageStruct{problems, result, cid, GetIslogin(r)})
	} else {
		result := []model.User{}
		c := session.DB("oj").C("user")
		c.Find(bson.M{"joinedcontest": cid}).Sort("-contesttotalaced", "contesttotaltime").All(&result)
		Render.HTML(w, http.StatusFound, "contestBoard", boardPageStruct{problems, result, cid, GetIslogin(r)})
	}
}
