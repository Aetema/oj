package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Miloas/oj/model"
)

//HandleAddContest : handle add contest action
func HandleAddContest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		//Format submit form datetime type to golang time: 2016-2-13T1:00 -> 2016-2-13 1:00:00
		contestStartTime := strings.Join(strings.Split(r.Form["contestStartTime"][0], "T"), " ") + ":00"
		contestProblems := strings.Fields(r.Form["contestProblems"][0])
		contestName := r.Form["contestName"][0]
		contestHowlong, _ := strconv.Atoi(r.Form["contestHowlong"][0])
		contestDescription := r.Form["contestDescription"][0]
		session := getMongoS()
		c := session.DB("oj").C("contests")
		count, _ := c.Count()
		contestID := "1000" + strconv.Itoa(count)
		c.Insert(&model.Contest{contestID, contestName, contestDescription, contestStartTime, contestHowlong, contestProblems})
	}
}

//HandleRemoveContest : handle remove contest action
func HandleRemoveContest(w http.ResponseWriter, r *http.Request) {

}

//HandleUpdateContest : handle update contest action
func HandleUpdateContest(w http.ResponseWriter, r *http.Request) {

}
