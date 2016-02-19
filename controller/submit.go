package controller

import (
	"encoding/json"
	"net/http"
)

type judgeQueueNode struct {
	User string
	ID   string
	Code string
}

//HandleSubmitCode : handle submited code action
func HandleSubmitCode(w http.ResponseWriter, r *http.Request) {
	c := RedisPool.Get()
	defer c.Close()
	r.ParseForm()
	user := getLoginUser(r)
	//c.Do("LPUSH", "judgeQueue", r.Form["submitedCode"][0])
	sendData, _ := json.Marshal(&judgeQueueNode{
		User: user.Username,
		ID:   r.URL.Query().Get("id"),
		Code: r.Form["submitedCode"][0]})
	c.Do("LPUSH", "judgeQueue", sendData)
	http.Redirect(w, r, "/status", http.StatusFound)
}
