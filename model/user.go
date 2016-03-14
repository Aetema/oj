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
	FirstBlood        []bool
}

//ClosedUser : when board close , we only display the situation an hour ago
type ClosedUser struct {
	Username string

	JoinedContest       string
	ContestTotalAced    int
	ContestTotalTime    int
	ContestAcedProblems []string
	ContestWrongTimes   []int
	ContestAcedTime     []int
	FirstBlood          []bool
}
