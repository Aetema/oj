package controller

import (
	"net/http"
	"strconv"

	"gopkg.in/boj/redistore.v1"

	"github.com/Miloas/oj/model"
	"github.com/gorilla/context"
)

const problemsPageNum int = 1

type problemsPageStruct struct {
	CurrentPage  int
	NextPage     int
	PreviousPage int
	CanNext      bool
	CanPrevious  bool
	Pagination   []int
	Problems     []model.Problem
	Islogin      bool
	Isadmin      bool

	//index info content
	HaveInfo bool
	Info     string
}

//HandleHome :handle "/"
func HandleHome(w http.ResponseWriter, r *http.Request) {
	//假设page是整数,回头改这
	p := 0
	if tmp := r.URL.Query().Get("page"); tmp != "" {
		p, _ = strconv.Atoi(tmp)
	}
	session := getMongoS()
	defer session.Close()
	c := session.DB("oj").C("problems")
	count, err := c.Count()
	totalPage := (count + problemsPageNum - 1) / problemsPageNum
	problems := []model.Problem{}
	err = c.Find(nil).Limit(problemsPageNum).Skip(problemsPageNum * p).All(&problems)
	if err != nil {
		panic(err)
	}
	pagination := []int{}
	for i := 0; i < totalPage; i++ {
		pagination = append(pagination, i)
	}
	canNext, canPrevious := false, false
	if p+1 < totalPage {
		canNext = true
	}
	if p-1 >= 0 {
		canPrevious = true
	}
	store, err := redistore.NewRediStore(10, "tcp", ":6379", "", []byte("secret-key"))
	if err != nil {
		panic(err)
	}
	defer store.Close()
	// Get a session.
	accountSession, _ := store.Get(r, "info")
	// Add a value.
	loginInfo := accountSession.Values["loginInfo"]
	// Save.
	accountSession.Options.MaxAge = -1
	accountSession.Save(r, w)
	ok := true
	val := ""
	if loginInfo == nil {
		ok = false
	} else {
		val = loginInfo.(string)
	}
	result := problemsPageStruct{p, p + 1, p - 1, canNext, canPrevious, pagination, problems, getIslogin(r), getIsadmin(r), ok, val}
	//defer store.Close()
	context.Clear(r)
	Render.HTML(w, http.StatusOK, "index", result)
}
