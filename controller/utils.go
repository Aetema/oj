package controller

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"

	"github.com/Miloas/oj/model"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/boj/redistore.v1"
	"gopkg.in/mgo.v2"
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
