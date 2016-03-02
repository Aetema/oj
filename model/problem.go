package model

//Problem : problem model
type Problem struct {
	ID             string
	Title          string
	Time           int
	Memory         int
	Description    string
	SampleInput    string
	SampleOutput   string
	StandardInput  string
	StandardOutput string
	Display        int
	Ac             int
	Total          int
	AcRate         int
}
