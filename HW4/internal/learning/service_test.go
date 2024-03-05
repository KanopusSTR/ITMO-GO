package learning

import (
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"sort"
	"testing"
)

func TestService_GetTutorsIDPreferIndividual(t *testing.T) {
	t.Helper()
	testCasesStudents := []struct {
		testName string
		name     string
		id       int64
		age      int64
		Subject  string
	}{
		{"first", "Dima", 0, 15, "Math"},
		{"second", "Moahim", 1, 20, "PE"},
		{"third", "Gwyndolin", 2, 30, "Art"},
	}

	testCasesIndividual := map[string][]int64{"Math": {0, 1, 2, 3, 4}}
	testCasesGroup := map[string][]int64{"Math": {5}, "PE": {6, 7, 8}}

	mc := minimock.NewController(t)
	defer mc.Finish()

	for _, tc := range testCasesStudents {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()

			repoMock := NewRepoMock(mc)
			repoIndividualMock := NewRepoIndividualMock(mc)
			repoGroupMock := NewRepoGroupMock(mc)

			retValue, ok := testCasesIndividual[tc.Subject]
			if !ok {
				retValue, ok = testCasesGroup[tc.Subject]
				if !ok {
					retValue = nil
				}
			}

			repoMock.GetStudentInfoMock.Expect(tc.id).Return(&studentInfo{tc.name, tc.age, tc.Subject}, true)
			repoIndividualMock.TutorsIDMock.Expect(tc.Subject).Return(retValue)
			repoGroupMock.TutorsIDMock.Expect(tc.Subject).Return(retValue)

			service := NewService(repoIndividualMock, repoGroupMock, repoMock)
			gotResult, ans := service.GetTutorsIDPreferIndividual(tc.id)

			require.Equal(t, retValue, gotResult)
			require.Equal(t, ok, ans)
		})
	}
}

func TestService_GetTopSubjects(t *testing.T) {
	t.Helper()
	testCasesSubjects := []struct {
		testName string
		subjs    []subjectInfo
		topN     int
		ok       bool
	}{
		{"test1", []subjectInfo{
			{"Math", 0},
			{"PE", 5},
			{"Art", 10}}, 2, true},
		{"test2", []subjectInfo{{"Math", 0}}, 0, true},
		{"test3", []subjectInfo{
			{"Math", 0},
			{"PE", 5},
			{"Art", 10}}, 3, true},
		{"test4", []subjectInfo{}, 0, true},
		{"test5", []subjectInfo{{"Math", 0}}, 2, false},
	}

	mc := minimock.NewController(t)
	defer mc.Finish()

	for _, tc := range testCasesSubjects {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()

			repoMock := NewRepoMock(mc)
			repoIndividualMock := NewRepoIndividualMock(mc)
			repoGroupMock := NewRepoGroupMock(mc)

			sort.SliceStable(tc.subjs, func(i, j int) bool {
				return tc.subjs[i].numberOfTutors < tc.subjs[j].numberOfTutors
			})

			var retValue []string

			if tc.topN == 0 {
				retValue = []string{}
			} else if tc.topN <= len(tc.subjs) {
				retValue = fromSubject(tc.subjs[:tc.topN])
			} else {
				retValue = nil
			}

			repoMock.GetAllSubjectsInfoMock.Expect().Return(tc.subjs, true)

			service := NewService(repoIndividualMock, repoGroupMock, repoMock)
			gotResult, ans := service.GetTopSubjects(tc.topN)

			require.Equal(t, retValue, gotResult)
			require.Equal(t, tc.ok, ans)
		})
	}
}
