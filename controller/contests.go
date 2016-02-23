package controller

import (
	"net/http"

	"github.com/Miloas/oj/model"
)

type contestPageStruct struct {
	CurrentPage  int
	NextPage     int
	PreviousPage int
	CanNext      bool
	CanPrevious  bool
	Status       []model.Contest
	Islogin      bool
}

//HandleContests : handle contests page
func HandleContests(w http.ResponseWriter, r *http.Request) {
	Render.HTML(w, http.StatusOK, "contests", nil)
}
