package model

//Status :status model
type Status struct {
	Sid        int
	User       string
	ID         string
	SubmitTime string
	Time       string
	Memory     string
	Result     string
	Lang       string
	ContestID  string

	Display bool
}
