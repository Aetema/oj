package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Miloas/oj/model"
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

		session := getMongoS()
		defer session.Close()
		statusCol := session.DB("oj").C("status")
		count, err := statusCol.Count()
		if err != nil {
			panic(err)
		}
		statusCol.Insert(&model.Status{count, user.Username, pid, "-", "-", "-", "Queue", r.Form["lang"][0], cid})
		//c.Do("LPUSH", "judgeQueue", r.Form["submitedCode"][0])
		sendData, _ := json.Marshal(&judgeQueueNode{
			Sid:       count,
			User:      user.Username,
			ID:        pid,
			Code:      r.Form["submitedCode"][0],
			Lang:      r.Form["lang"][0],
			ContestID: cid})
		c.Do("LPUSH", "judgeQueue", sendData)
		http.Redirect(w, r, "/contest/status?cid="+cid, http.StatusFound)
	}
}
