package model

//User : user model
type User struct {
	Username string
	Password string
	Role     string
	Accepted []string

	//contest info (include contest grade) , should use nested struct , here is a dirty hack, improve it later
	JoinedContest       string
	ContestTotalAced    int
	ContestTotalTime    int
	ContestAcedProblems []string
	//record 每道题的错误次数
	ContestWrongTimes []int
	ContestAcedTime   []int
}
