package controller

import (
	"net/http"
	"strconv"

	"github.com/Miloas/oj/model"
)

type contestPageStruct struct {
	CurrentPage  int
	NextPage     int
	PreviousPage int
	CanNext      bool
	CanPrevious  bool
	Contests     []model.Contest
	Islogin      bool
	IsAdmin      bool
}

const contestsPageNum int = 5

//HandleContests : handle contests page
func HandleContests(w http.ResponseWriter, r *http.Request) {
	p := 0
	if tmp := r.URL.Query().Get("page"); tmp != "" {
		p, _ = strconv.Atoi(tmp)
	}
	session := getMongoS()
	defer session.Close()
	c := session.DB("oj").C("contests")
	count, err := c.Count()
	totalPage := (count + contestsPageNum - 1) / contestsPageNum
	contests := []model.Contest{}
	err = c.Find(nil).Sort("-starttime").Limit(contestsPageNum).Skip(contestsPageNum * p).All(&contests)
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
	result := contestPageStruct{p, p + 1, p - 1, canNext, canPrevious, contests, GetIslogin(r), GetIsadmin(r)}
	Render.HTML(w, http.StatusOK, "contests", result)
}
