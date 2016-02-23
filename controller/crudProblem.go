package controller

import (
	"net/http"
	"strconv"

	"github.com/Miloas/oj/model"

	"gopkg.in/mgo.v2/bson"
)

//HandleAddProblem : handle add problem request
func HandleAddProblem(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		session := getMongoS()
		defer session.Close()
		c := session.DB("oj").C("problems")
		count, err := c.Count()
		if err != nil {
			count = 0
		}
		t, _ := strconv.Atoi(r.Form["time"][0])
		m, _ := strconv.Atoi(r.Form["memory"][0])
		newProblem := &model.Problem{strconv.Itoa(count + 1000), r.Form["title"][0], t,
			m, r.Form["description"][0], r.Form["sampleinput"][0],
			r.Form["sampleoutput"][0], r.Form["standardinput"][0], r.Form["standardoutput"][0],
			0, 0, 0, 0}
		err = c.Insert(newProblem)
		if err != nil {
			panic(err)
		}
	}
}

//HandleRemoveProblem : handle remove problem request
func HandleRemoveProblem(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		session := getMongoS()
		defer session.Close()
		c := session.DB("oj").C("problems")
		err := c.Remove(bson.M{"id": r.Form["id"][0]})
		if err != nil {
			panic(err)
		}
	}
}

//HandleUpdateProblem : handle Update problem request
func HandleUpdateProblem(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		session := getMongoS()
		defer session.Close()
		c := session.DB("oj").C("problems")
		t, _ := strconv.Atoi(r.Form["time"][0])
		m, _ := strconv.Atoi(r.Form["memory"][0])
		display, _ := strconv.Atoi(r.Form["display"][0])
		err := c.Update(bson.M{"id": r.Form["id"][0]},
			bson.M{"$set": bson.M{
				"title":          r.Form["title"][0],
				"time":           t,
				"memory":         m,
				"description":    r.Form["description"][0],
				"sampleinput":    r.Form["sampleinput"][0],
				"sampleoutput":   r.Form["sampleoutput"][0],
				"standardinput":  r.Form["standardinput"][0],
				"standardoutput": r.Form["standardoutput"][0],
				"display":        display,
				"ac":             0,
				"total":          0,
				"acrate":         0,
			}})
		if err != nil {
			panic(err)
		}
	}
}
