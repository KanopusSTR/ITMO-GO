package learning

type Repo interface {
	GetStudentInfo(id int64) (*studentInfo, bool)
	GetAllSubjects() ([]string, bool)
	GetAllSubjectsInfo() ([]subjectInfo, bool)
}
