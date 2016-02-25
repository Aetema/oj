package controller

import (
	"encoding/json"
	"net/http"
)

type judgeQueueNode struct {
	User string
	ID   string
	Code string

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
		//c.Do("LPUSH", "judgeQueue", r.Form["submitedCode"][0])
		sendData, _ := json.Marshal(&judgeQueueNode{
			User:      user.Username,
			ID:        r.URL.Query().Get("id"),
			Code:      r.Form["submitedCode"][0],
			ContestID: ""})
		c.Do("LPUSH", "judgeQueue", sendData)
		http.Redirect(w, r, "/status", http.StatusFound)
	}
}
