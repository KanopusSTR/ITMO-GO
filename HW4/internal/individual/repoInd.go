package individual

type RepoIndividual interface {
	TutorsID(subject string) []int64
	Subjects() []string
}
