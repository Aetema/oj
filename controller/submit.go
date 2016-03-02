package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Miloas/oj/model"
)

type judgeQueueNode struct {
	Sid  int
	User string
	ID   string
	Code string
	Lang string

	//contest Info
	ContestID string
}

//HandleSubmitCode : handle submited code action
func HandleSubmitCode(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		c := RedisPool.Get()
		defer c.Close()
		r.ParseForm()
		user := GetLoginUser(r)
		session := getMongoS()
		defer session.Close()
		statusCol := session.DB("oj").C("status")
		count, err := statusCol.Count()
		if err != nil {
			panic(err)
		}
		display := true
		if GetIsadmin(r) {
			display = false
		}
		statusCol.Insert(&model.Status{count, user.Username, r.URL.Query().Get("id"), "-", "-", "-", "Queue", r.Form["lang"][0], "", display})
		//c.Do("LPUSH", "judgeQueue", r.Form["submitedCode"][0])
		sendData, _ := json.Marshal(&judgeQueueNode{
			Sid:       count,
			User:      user.Username,
			ID:        r.URL.Query().Get("id"),
			Code:      r.Form["submitedCode"][0],
			Lang:      r.Form["lang"][0],
			ContestID: ""})
		c.Do("LPUSH", "judgeQueue", sendData)
		http.Redirect(w, r, "/status", http.StatusFound)
	}
}
