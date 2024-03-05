package group

type RepoGroup interface {
	TutorsID(subject string) []int64
	Subjects() []string
}
