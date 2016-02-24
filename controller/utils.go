package controller

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Miloas/oj/model"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/boj/redistore.v1"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func getMongoS() *mgo.Session {
	if S == nil {
		var err error
		S, err = mgo.Dial("localhost:27017")
		if err != nil {
			panic(err)
		}
	}
	return S.Clone()
}

func newRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "localhost:6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

func cryptoPassword(password string) string {
	h := md5.New()
	h.Write([]byte(password))
	return hex.EncodeToString(h.Sum(nil))
}

//GetIslogin : check if or not login
func GetIslogin(r *http.Request) bool {
	store, err := redistore.NewRediStore(10, "tcp", ":6379", "", []byte("secret-key"))
	if err != nil {
		panic(err)
	}
	defer store.Close()
	accountSession, _ := store.Get(r, "user")
	// fmt.Println(accountSession.Values["currentuser"])
	return accountSession.Values["currentuser"] != nil
}

//GetLoginUser : get login user info
func GetLoginUser(r *http.Request) *model.User {
	store, err := redistore.NewRediStore(10, "tcp", ":6379", "", []byte("secret-key"))
	if err != nil {
		panic(err)
	}
	defer store.Close()
	accountSession, _ := store.Get(r, "user")
	return accountSession.Values["currentuser"].(*model.User)
}

//GetIsadmin : check if or not login user is admin
func GetIsadmin(r *http.Request) bool {
	if !GetIslogin(r) {
		return false
	}
	store, err := redistore.NewRediStore(10, "tcp", ":6379", "", []byte("secret-key"))
	if err != nil {
		panic(err)
	}
	defer store.Close()
	accountSession, _ := store.Get(r, "user")
	return accountSession.Values["currentuser"].(*model.User).Role == "admin"
}

//CheckInStringArray : check key if or not in stringArray
func CheckInStringArray(key string, stringArray []string) bool {
	if len(stringArray) <= 0 {
		return false
	}
	for _, val := range stringArray {
		if val == key {
			return true
		}
	}
	return false
}

//FormDatetime2Gotime : Format submit form datetime type to golang time: 2016-2-13T1:00 -> 2016-2-13 1:00:00
func FormDatetime2Gotime(x string) string {
	return strings.Join(strings.Split(x, "T"), " ") + ":00"
}

//CheckAuth2Problem : check auth to /problem?id=:id and /problem/submit?id=:id , normal user only can touch display==1 problem , admin user can touch everything
func CheckAuth2Problem(r *http.Request) bool {
	id := r.URL.Query().Get("id")
	session := getMongoS()
	defer session.Close()
	c := session.DB("oj").C("problems")
	result := []model.Problem{}
	err := c.Find(bson.M{"id": id}).All(&result)
	if err != nil {
		panic(err)
	}
	if len(result) > 0 && result[0].Display == 1 {
		return true
	}
	if len(result) > 0 && GetIsadmin(r) {
		return true
	}
	return false
}

//CheckAuth2Contest : check auth to /contest /contest/problems /contest/problems/submit /contest/board /contest/status, only admin and joined user (during contest) can touch
func CheckAuth2Contest(r *http.Request) bool {
	cid := r.URL.Query().Get("cid")
	if cid == "" {
		return false
	}
	if GetIsadmin(r) {
		return true
	}
	if !GetIslogin(r) {
		return false
	}
	user := GetLoginUser(r)
	if user.JoinedContest != cid {
		return false
	}
	session := getMongoS()
	defer session.Close()
	contestCol := session.DB("oj").C("contests")
	contest := []model.Contest{}
	contestCol.Find(bson.M{"contestid": cid}).All(&contest)
	if len(contest) <= 0 {
		return false
	}
	//2016-2-24 10:25:00
	contestStartTime, _ := time.Parse("2006-01-02 15:04:05", contest[0].StartTime)
	contestHowlong := contest[0].HowLong
	contestEndTime := contestStartTime
	for i := 0; i < contestHowlong; i++ {
		contestEndTime = contestEndTime.Add(time.Hour)
	}
	currentTime, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(contestStartTime, contestEndTime, currentTime)
	if currentTime.After(contestEndTime) || currentTime.Before(contestStartTime) {
		return false
	}
	return true
}
