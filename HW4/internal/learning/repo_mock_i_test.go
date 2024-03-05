package learning

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i hw4/internal/individual.RepoIndividual -o ../learning/repo_mock_i_test.go -n RepoIndividualMock

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// RepoIndividualMock implements individual.RepoIndividual
type RepoIndividualMock struct {
	t minimock.Tester

	funcSubjects          func() (sa1 []string)
	inspectFuncSubjects   func()
	afterSubjectsCounter  uint64
	beforeSubjectsCounter uint64
	SubjectsMock          mRepoIndividualMockSubjects

	funcTutorsID          func(subject string) (ia1 []int64)
	inspectFuncTutorsID   func(subject string)
	afterTutorsIDCounter  uint64
	beforeTutorsIDCounter uint64
	TutorsIDMock          mRepoIndividualMockTutorsID
}

// NewRepoIndividualMock returns a mock for individual.RepoIndividual
func NewRepoIndividualMock(t minimock.Tester) *RepoIndividualMock {
	m := &RepoIndividualMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.SubjectsMock = mRepoIndividualMockSubjects{mock: m}

	m.TutorsIDMock = mRepoIndividualMockTutorsID{mock: m}
	m.TutorsIDMock.callArgs = []*RepoIndividualMockTutorsIDParams{}

	return m
}

type mRepoIndividualMockSubjects struct {
	mock               *RepoIndividualMock
	defaultExpectation *RepoIndividualMockSubjectsExpectation
	expectations       []*RepoIndividualMockSubjectsExpectation
}

// RepoIndividualMockSubjectsExpectation specifies expectation struct of the RepoIndividual.Subjects
type RepoIndividualMockSubjectsExpectation struct {
	mock *RepoIndividualMock

	results *RepoIndividualMockSubjectsResults
	Counter uint64
}

// RepoIndividualMockSubjectsResults contains results of the RepoIndividual.Subjects
type RepoIndividualMockSubjectsResults struct {
	sa1 []string
}

// Expect sets up expected params for RepoIndividual.Subjects
func (mmSubjects *mRepoIndividualMockSubjects) Expect() *mRepoIndividualMockSubjects {
	if mmSubjects.mock.funcSubjects != nil {
		mmSubjects.mock.t.Fatalf("RepoIndividualMock.Subjects mock is already set by Set")
	}

	if mmSubjects.defaultExpectation == nil {
		mmSubjects.defaultExpectation = &RepoIndividualMockSubjectsExpectation{}
	}

	return mmSubjects
}

// Inspect accepts an inspector function that has same arguments as the RepoIndividual.Subjects
func (mmSubjects *mRepoIndividualMockSubjects) Inspect(f func()) *mRepoIndividualMockSubjects {
	if mmSubjects.mock.inspectFuncSubjects != nil {
		mmSubjects.mock.t.Fatalf("Inspect function is already set for RepoIndividualMock.Subjects")
	}

	mmSubjects.mock.inspectFuncSubjects = f

	return mmSubjects
}

// Return sets up results that will be returned by RepoIndividual.Subjects
func (mmSubjects *mRepoIndividualMockSubjects) Return(sa1 []string) *RepoIndividualMock {
	if mmSubjects.mock.funcSubjects != nil {
		mmSubjects.mock.t.Fatalf("RepoIndividualMock.Subjects mock is already set by Set")
	}

	if mmSubjects.defaultExpectation == nil {
		mmSubjects.defaultExpectation = &RepoIndividualMockSubjectsExpectation{mock: mmSubjects.mock}
	}
	mmSubjects.defaultExpectation.results = &RepoIndividualMockSubjectsResults{sa1}
	return mmSubjects.mock
}

// Set uses given function f to mock the RepoIndividual.Subjects method
func (mmSubjects *mRepoIndividualMockSubjects) Set(f func() (sa1 []string)) *RepoIndividualMock {
	if mmSubjects.defaultExpectation != nil {
		mmSubjects.mock.t.Fatalf("Default expectation is already set for the RepoIndividual.Subjects method")
	}

	if len(mmSubjects.expectations) > 0 {
		mmSubjects.mock.t.Fatalf("Some expectations are already set for the RepoIndividual.Subjects method")
	}

	mmSubjects.mock.funcSubjects = f
	return mmSubjects.mock
}

// Subjects implements individual.RepoIndividual
func (mmSubjects *RepoIndividualMock) Subjects() (sa1 []string) {
	mm_atomic.AddUint64(&mmSubjects.beforeSubjectsCounter, 1)
	defer mm_atomic.AddUint64(&mmSubjects.afterSubjectsCounter, 1)

	if mmSubjects.inspectFuncSubjects != nil {
		mmSubjects.inspectFuncSubjects()
	}

	if mmSubjects.SubjectsMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSubjects.SubjectsMock.defaultExpectation.Counter, 1)

		mm_results := mmSubjects.SubjectsMock.defaultExpectation.results
		if mm_results == nil {
			mmSubjects.t.Fatal("No results are set for the RepoIndividualMock.Subjects")
		}
		return (*mm_results).sa1
	}
	if mmSubjects.funcSubjects != nil {
		return mmSubjects.funcSubjects()
	}
	mmSubjects.t.Fatalf("Unexpected call to RepoIndividualMock.Subjects.")
	return
}

// SubjectsAfterCounter returns a count of finished RepoIndividualMock.Subjects invocations
func (mmSubjects *RepoIndividualMock) SubjectsAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSubjects.afterSubjectsCounter)
}

// SubjectsBeforeCounter returns a count of RepoIndividualMock.Subjects invocations
func (mmSubjects *RepoIndividualMock) SubjectsBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSubjects.beforeSubjectsCounter)
}

// MinimockSubjectsDone returns true if the count of the Subjects invocations corresponds
// the number of defined expectations
func (m *RepoIndividualMock) MinimockSubjectsDone() bool {
	for _, e := range m.SubjectsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SubjectsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSubjectsCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSubjects != nil && mm_atomic.LoadUint64(&m.afterSubjectsCounter) < 1 {
		return false
	}
	return true
}

// MinimockSubjectsInspect logs each unmet expectation
func (m *RepoIndividualMock) MinimockSubjectsInspect() {
	for _, e := range m.SubjectsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to RepoIndividualMock.Subjects")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SubjectsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSubjectsCounter) < 1 {
		m.t.Error("Expected call to RepoIndividualMock.Subjects")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSubjects != nil && mm_atomic.LoadUint64(&m.afterSubjectsCounter) < 1 {
		m.t.Error("Expected call to RepoIndividualMock.Subjects")
	}
}

type mRepoIndividualMockTutorsID struct {
	mock               *RepoIndividualMock
	defaultExpectation *RepoIndividualMockTutorsIDExpectation
	expectations       []*RepoIndividualMockTutorsIDExpectation

	callArgs []*RepoIndividualMockTutorsIDParams
	mutex    sync.RWMutex
}

// RepoIndividualMockTutorsIDExpectation specifies expectation struct of the RepoIndividual.TutorsID
type RepoIndividualMockTutorsIDExpectation struct {
	mock    *RepoIndividualMock
	params  *RepoIndividualMockTutorsIDParams
	results *RepoIndividualMockTutorsIDResults
	Counter uint64
}

// RepoIndividualMockTutorsIDParams contains parameters of the RepoIndividual.TutorsID
type RepoIndividualMockTutorsIDParams struct {
	subject string
}

// RepoIndividualMockTutorsIDResults contains results of the RepoIndividual.TutorsID
type RepoIndividualMockTutorsIDResults struct {
	ia1 []int64
}

// Expect sets up expected params for RepoIndividual.TutorsID
func (mmTutorsID *mRepoIndividualMockTutorsID) Expect(subject string) *mRepoIndividualMockTutorsID {
	if mmTutorsID.mock.funcTutorsID != nil {
		mmTutorsID.mock.t.Fatalf("RepoIndividualMock.TutorsID mock is already set by Set")
	}

	if mmTutorsID.defaultExpectation == nil {
		mmTutorsID.defaultExpectation = &RepoIndividualMockTutorsIDExpectation{}
	}

	mmTutorsID.defaultExpectation.params = &RepoIndividualMockTutorsIDParams{subject}
	for _, e := range mmTutorsID.expectations {
		if minimock.Equal(e.params, mmTutorsID.defaultExpectation.params) {
			mmTutorsID.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmTutorsID.defaultExpectation.params)
		}
	}

	return mmTutorsID
}

// Inspect accepts an inspector function that has same arguments as the RepoIndividual.TutorsID
func (mmTutorsID *mRepoIndividualMockTutorsID) Inspect(f func(subject string)) *mRepoIndividualMockTutorsID {
	if mmTutorsID.mock.inspectFuncTutorsID != nil {
		mmTutorsID.mock.t.Fatalf("Inspect function is already set for RepoIndividualMock.TutorsID")
	}

	mmTutorsID.mock.inspectFuncTutorsID = f

	return mmTutorsID
}

// Return sets up results that will be returned by RepoIndividual.TutorsID
func (mmTutorsID *mRepoIndividualMockTutorsID) Return(ia1 []int64) *RepoIndividualMock {
	if mmTutorsID.mock.funcTutorsID != nil {
		mmTutorsID.mock.t.Fatalf("RepoIndividualMock.TutorsID mock is already set by Set")
	}

	if mmTutorsID.defaultExpectation == nil {
		mmTutorsID.defaultExpectation = &RepoIndividualMockTutorsIDExpectation{mock: mmTutorsID.mock}
	}
	mmTutorsID.defaultExpectation.results = &RepoIndividualMockTutorsIDResults{ia1}
	return mmTutorsID.mock
}

// Set uses given function f to mock the RepoIndividual.TutorsID method
func (mmTutorsID *mRepoIndividualMockTutorsID) Set(f func(subject string) (ia1 []int64)) *RepoIndividualMock {
	if mmTutorsID.defaultExpectation != nil {
		mmTutorsID.mock.t.Fatalf("Default expectation is already set for the RepoIndividual.TutorsID method")
	}

	if len(mmTutorsID.expectations) > 0 {
		mmTutorsID.mock.t.Fatalf("Some expectations are already set for the RepoIndividual.TutorsID method")
	}

	mmTutorsID.mock.funcTutorsID = f
	return mmTutorsID.mock
}

// When sets expectation for the RepoIndividual.TutorsID which will trigger the result defined by the following
// Then helper
func (mmTutorsID *mRepoIndividualMockTutorsID) When(subject string) *RepoIndividualMockTutorsIDExpectation {
	if mmTutorsID.mock.funcTutorsID != nil {
		mmTutorsID.mock.t.Fatalf("RepoIndividualMock.TutorsID mock is already set by Set")
	}

	expectation := &RepoIndividualMockTutorsIDExpectation{
		mock:   mmTutorsID.mock,
		params: &RepoIndividualMockTutorsIDParams{subject},
	}
	mmTutorsID.expectations = append(mmTutorsID.expectations, expectation)
	return expectation
}

// Then sets up RepoIndividual.TutorsID return parameters for the expectation previously defined by the When method
func (e *RepoIndividualMockTutorsIDExpectation) Then(ia1 []int64) *RepoIndividualMock {
	e.results = &RepoIndividualMockTutorsIDResults{ia1}
	return e.mock
}

// TutorsID implements individual.RepoIndividual
func (mmTutorsID *RepoIndividualMock) TutorsID(subject string) (ia1 []int64) {
	mm_atomic.AddUint64(&mmTutorsID.beforeTutorsIDCounter, 1)
	defer mm_atomic.AddUint64(&mmTutorsID.afterTutorsIDCounter, 1)

	if mmTutorsID.inspectFuncTutorsID != nil {
		mmTutorsID.inspectFuncTutorsID(subject)
	}

	mm_params := &RepoIndividualMockTutorsIDParams{subject}

	// Record call args
	mmTutorsID.TutorsIDMock.mutex.Lock()
	mmTutorsID.TutorsIDMock.callArgs = append(mmTutorsID.TutorsIDMock.callArgs, mm_params)
	mmTutorsID.TutorsIDMock.mutex.Unlock()

	for _, e := range mmTutorsID.TutorsIDMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.ia1
		}
	}

	if mmTutorsID.TutorsIDMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmTutorsID.TutorsIDMock.defaultExpectation.Counter, 1)
		mm_want := mmTutorsID.TutorsIDMock.defaultExpectation.params
		mm_got := RepoIndividualMockTutorsIDParams{subject}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmTutorsID.t.Errorf("RepoIndividualMock.TutorsID got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmTutorsID.TutorsIDMock.defaultExpectation.results
		if mm_results == nil {
			mmTutorsID.t.Fatal("No results are set for the RepoIndividualMock.TutorsID")
		}
		return (*mm_results).ia1
	}
	if mmTutorsID.funcTutorsID != nil {
		return mmTutorsID.funcTutorsID(subject)
	}
	mmTutorsID.t.Fatalf("Unexpected call to RepoIndividualMock.TutorsID. %v", subject)
	return
}

// TutorsIDAfterCounter returns a count of finished RepoIndividualMock.TutorsID invocations
func (mmTutorsID *RepoIndividualMock) TutorsIDAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmTutorsID.afterTutorsIDCounter)
}

// TutorsIDBeforeCounter returns a count of RepoIndividualMock.TutorsID invocations
func (mmTutorsID *RepoIndividualMock) TutorsIDBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmTutorsID.beforeTutorsIDCounter)
}

// Calls returns a list of arguments used in each call to RepoIndividualMock.TutorsID.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmTutorsID *mRepoIndividualMockTutorsID) Calls() []*RepoIndividualMockTutorsIDParams {
	mmTutorsID.mutex.RLock()

	argCopy := make([]*RepoIndividualMockTutorsIDParams, len(mmTutorsID.callArgs))
	copy(argCopy, mmTutorsID.callArgs)

	mmTutorsID.mutex.RUnlock()

	return argCopy
}

// MinimockTutorsIDDone returns true if the count of the TutorsID invocations corresponds
// the number of defined expectations
func (m *RepoIndividualMock) MinimockTutorsIDDone() bool {
	for _, e := range m.TutorsIDMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.TutorsIDMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterTutorsIDCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcTutorsID != nil && mm_atomic.LoadUint64(&m.afterTutorsIDCounter) < 1 {
		return false
	}
	return true
}

// MinimockTutorsIDInspect logs each unmet expectation
func (m *RepoIndividualMock) MinimockTutorsIDInspect() {
	for _, e := range m.TutorsIDMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RepoIndividualMock.TutorsID with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.TutorsIDMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterTutorsIDCounter) < 1 {
		if m.TutorsIDMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RepoIndividualMock.TutorsID")
		} else {
			m.t.Errorf("Expected call to RepoIndividualMock.TutorsID with params: %#v", *m.TutorsIDMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcTutorsID != nil && mm_atomic.LoadUint64(&m.afterTutorsIDCounter) < 1 {
		m.t.Error("Expected call to RepoIndividualMock.TutorsID")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *RepoIndividualMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockSubjectsInspect()

		m.MinimockTutorsIDInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *RepoIndividualMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *RepoIndividualMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockSubjectsDone() &&
		m.MinimockTutorsIDDone()
}