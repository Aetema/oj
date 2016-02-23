package model

//Contest : contest model
type Contest struct {
	ContestID          string
	ContestName        string
	ContestDescription string
	StartTime          string
	HowLong            int //hours
	ContestProblems    []string
}
