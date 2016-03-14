package model

//Contest : contest model
type Contest struct {
	ContestID           string
	ContestName         string
	ContestDescription  string
	FormStartTime       string
	StartTime           string
	HowLong             int //hours
	FormContestProblems string
	ContestProblems     []string

	HaveAced []bool
}
