package controller

import (
	"encoding/json"
	"net/http"
)

//HandleContestSubmit : handle contest submit action
func HandleContestSubmit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		cid := r.URL.Query().Get("cid")
		pid := r.URL.Query().Get("pid")
		c := RedisPool.Get()
		defer c.Close()
		r.ParseForm()
		user := GetLoginUser(r)
		//c.Do("LPUSH", "judgeQueue", r.Form["submitedCode"][0])
		sendData, _ := json.Marshal(&judgeQueueNode{
			User:      user.Username,
			ID:        pid,
			Code:      r.Form["submitedCode"][0],
			ContestID: cid})
		c.Do("LPUSH", "judgeQueue", sendData)
		http.Redirect(w, r, "/contest/status?cid="+cid, http.StatusFound)
	}
}
