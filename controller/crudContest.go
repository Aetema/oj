package controller

import (
	"net/http"
	"strconv"
	"strings"

	"gopkg.in/mgo.v2/bson"

	"github.com/Miloas/oj/model"
)

//HandleAddContest : handle add contest action
func HandleAddContest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		contestStartTime := FormDatetime2Gotime(r.Form["contestStartTime"][0])
		contestProblems := strings.Fields(r.Form["contestProblems"][0])
		contestName := r.Form["contestName"][0]
		contestHowlong, _ := strconv.Atoi(r.Form["contestHowlong"][0])
		contestDescription := r.Form["contestDescription"][0]
		session := getMongoS()
		defer session.Close()
		c := session.DB("oj").C("contests")
		count, _ := c.Count()
		contestID := strconv.Itoa(1000 + count)
		c.Insert(&model.Contest{contestID, contestName, contestDescription, r.Form["contestStartTime"][0], contestStartTime, contestHowlong, strings.Join(contestProblems, " "), contestProblems})
	}
}

//HandleRemoveContest : handle remove contest action
func HandleRemoveContest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		session := getMongoS()
		defer session.Close()
		c := session.DB("oj").C("contests")
		err := c.Remove(bson.M{"contestid": r.Form["contestID"][0]})
		if err != nil {
			panic(err)
		}
	}
}

//HandleUpdateContest : handle update contest action
func HandleUpdateContest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		session := getMongoS()
		defer session.Close()
		c := session.DB("oj").C("contests")
		contestHowlong, _ := strconv.Atoi(r.Form["contestHowlong"][0])
		err := c.Update(bson.M{"contestid": r.Form["contestID"][0]},
			bson.M{"$set": bson.M{
				"contestname":         r.Form["contestName"][0],
				"contestdescription":  r.Form["contestDescription"][0],
				"formstarttime":       r.Form["contestStartTime"][0],
				"starttime":           FormDatetime2Gotime(r.Form["contestStartTime"][0]),
				"howlong":             contestHowlong,
				"formcontestproblems": strings.Join(strings.Fields(r.Form["contestProblems"][0]), " "),
				"contestproblems":     strings.Fields(r.Form["contestProblems"][0]),
			}})
		if err != nil {
			panic(err)
		}
	}
}
