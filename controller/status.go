package controller

import (
	"net/http"
	"strconv"

	"github.com/Miloas/oj/model"
)

const statusPageNum int = 10

type statusPageStruct struct {
	CurrentPage  int
	NextPage     int
	PreviousPage int
	CanNext      bool
	CanPrevious  bool
	Status       []model.Status
	Islogin      bool
}

//HandleStatus : Handle display status page
func HandleStatus(w http.ResponseWriter, r *http.Request) {
	p := 0
	if tmp := r.URL.Query().Get("page"); tmp != "" {
		p, _ = strconv.Atoi(tmp)
	}
	session := getMongoS()
	defer session.Close()
	c := session.DB("oj").C("status")
	count, err := c.Count()
	totalPage := (count + statusPageNum - 1) / statusPageNum
	status := []model.Status{}
	err = c.Find(nil).Sort("-submittime").Limit(statusPageNum).Skip(statusPageNum * p).All(&status)
	if err != nil {
		panic(err)
	}
	canNext, canPrevious := false, false
	if p+1 < totalPage {
		canNext = true
	}
	if p-1 >= 0 {
		canPrevious = true
	}
	result := statusPageStruct{p, p + 1, p - 1, canNext, canPrevious, status, getIslogin(r)}
	Render.HTML(w, http.StatusOK, "status", result)

}
