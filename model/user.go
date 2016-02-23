package model

//User : user model
type User struct {
	Username      string
	Password      string
	Role          string
	Accepted      []string
	JoinedContest string
}
