package models

type Problem struct {
	Type  string
	Ref   string
	Meta  interface{}
	Cause ProblemCause
	Fix   *ProblemFix
}

type ProblemCause struct {
}

type ProblemFix struct {
}
