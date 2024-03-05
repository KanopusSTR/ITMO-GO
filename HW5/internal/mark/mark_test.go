package mark

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestStudents_SummaryByStudent(t *testing.T) {
	testCasesStudents := []struct {
		testName     string
		students     []Student
		sableStudent string
		ans          int
		result       bool
	}{
		{"testOne", []Student{{Name: "Dima", Mark: 5}}, "Dima", 5, true},

		{"testMany", []Student{
			{Name: "Moahim", Mark: 4},
			{Name: "Dima", Mark: 1},
			{Name: "Moahim", Mark: 5},
		}, "Moahim", 9, true},

		{"testOdd", []Student{{Name: "Gwyndolin", Mark: 3}}, "Dima", 0, false},
	}

	for _, tc := range testCasesStudents {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()
			students := Students{tc.students}
			a, b := students.SummaryByStudent(tc.sableStudent)
			require.Equal(t, a, tc.ans)
			require.Equal(t, b, tc.result)
		})
	}
}

func TestStudents_AverageByStudent(t *testing.T) {
	testCasesStudents := []struct {
		testName     string
		students     []Student
		sableStudent string
		ans          float32
		result       bool
	}{
		{"testOneMark", []Student{{Name: "Dima", Mark: 5}}, "Dima", 5, true},

		{"testManyMarks", []Student{
			{Name: "Dima", Mark: 4},
			{Name: "Dima", Mark: 4},
			{Name: "Dima", Mark: 2},
		}, "Dima", 3.33, true},

		{"testManyStudents", []Student{
			{Name: "Dima", Mark: 4},
			{Name: "Dima", Mark: 4},
			{Name: "Dima", Mark: 2},
			{Name: "Moahim", Mark: 5},
		}, "Dima", 3.33, true},

		{"testOddStudent", []Student{{Name: "Gwyndolin", Mark: 3}}, "Dima", 0, false},
	}

	for _, tc := range testCasesStudents {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()
			students := Students{tc.students}
			a, b := students.AverageByStudent(tc.sableStudent)
			require.Equal(t, tc.ans, a)
			require.Equal(t, tc.result, b)
		})
	}
}

func TestStudents_Students(t *testing.T) {
	testCasesStudents := []struct {
		testName string
		students []Student
		ans      []string
	}{
		{"testOne", []Student{{Name: "Dima", Mark: 5}}, []string{"Dima"}},

		{"testMany", []Student{
			{Name: "Dima", Mark: 4},
			{Name: "Dima", Mark: 4},
			{Name: "Dima", Mark: 2},
			{Name: "Moahim", Mark: 5},
			{Name: "Petya", Mark: 9},
			{Name: "Sirius", Mark: 2},
			{Name: "Sirius", Mark: 4},
		}, []string{"Dima", "Petya", "Sirius", "Moahim"}},

		{"testEmpty", []Student{}, []string{}},
	}

	for _, tc := range testCasesStudents {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()
			students := Students{tc.students}
			a := students.Students()
			require.Equal(t, tc.ans, a)
		})
	}
}

func TestStudents_Summary(t *testing.T) {
	testCasesStudents := []struct {
		testName string
		students []Student
		ans      int
	}{
		{"testOne", []Student{{Name: "Dima", Mark: 5}}, 5},

		{"testMany", []Student{
			{Name: "Dima", Mark: 4},
			{Name: "Dima", Mark: 4},
			{Name: "Dima", Mark: 2},
			{Name: "Moahim", Mark: 5},
		}, 15},

		{"testEmpty", []Student{}, 0},
	}

	for _, tc := range testCasesStudents {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()
			students := Students{tc.students}
			a := students.Summary()
			require.Equal(t, tc.ans, a)
		})
	}
}

func TestStudents_Median(t *testing.T) {
	testCasesStudents := []struct {
		testName string
		students []Student
		ans      int
	}{
		{"testOne", []Student{{Name: "Dima", Mark: 5}}, 5},

		{"testMany", []Student{
			{Name: "Dima", Mark: 4},
			{Name: "Dima", Mark: 4},
			{Name: "Dima", Mark: 2},
			{Name: "Moahim", Mark: 5},
		}, 4},

		{"testEmpty", []Student{}, 0},
	}

	for _, tc := range testCasesStudents {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()
			students := Students{tc.students}
			a := students.Median()
			require.Equal(t, tc.ans, a)
		})
	}
}

func TestStudents_MostFrequent(t *testing.T) {
	testCasesStudents := []struct {
		testName string
		students []Student
		ans      int
	}{
		{"testOne", []Student{{Name: "Dima", Mark: 5}}, 5},

		{"testMany", []Student{
			{Name: "Dima", Mark: 4},
			{Name: "Dima", Mark: 4},
			{Name: "Dima", Mark: 2},
			{Name: "Moahim", Mark: 5},
		}, 4},

		{"testEmpty", []Student{}, 0},
	}

	for _, tc := range testCasesStudents {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()
			students := Students{tc.students}
			a := students.MostFrequent()
			require.Equal(t, tc.ans, a)
		})
	}
}

func TestReadStudentsStatistic(t *testing.T) {
	testCasesStudents := []struct {
		testName string
		students []Student
		fileName string
		error    bool
	}{
		{"testExistingFile", []Student{
			{Name: "Примеров Пример Примерович", Mark: 5},
			{Name: "Примеров Пример Примерович", Mark: 6},
			{Name: "Примеров Пример", Mark: 7},
		}, "../../data/input_1.tsv", false},

		{"testNonExistingFile", []Student{
			{Name: "Примеров Пример Примерович", Mark: 5},
			{Name: "Примеров Пример Примерович", Mark: 6},
			{Name: "Примеров Пример", Mark: 7},
		}, "../../data/input_2.tsv", true},

		{"testBrokenFile", []Student{
			{Name: "Примеров Пример Примерович", Mark: 6},
			{Name: "Примеров Пример", Mark: 7},
		}, "../../data/broken_input_1.tsv", true},
	}
	for _, tc := range testCasesStudents {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()

			reader, _ := os.OpenFile(tc.fileName, os.O_RDONLY, 0666)
			defer func(reader *os.File) {
				err := reader.Close()
				if err != nil {
					require.True(t, tc.error)
				}
			}(reader)

			studentsStatistic, err := ReadStudentsStatistic(reader)
			if err != nil {
				require.True(t, tc.error)
				return
			}

			students := Students{tc.students}
			require.Equal(t, students.Students(), studentsStatistic.Students())
			require.Equal(t, students.Summary(), studentsStatistic.Summary())
		})
	}
}

func TestWriteStudentsStatistic(t *testing.T) {
	testCasesStudents := []struct {
		testName string
		students []Student
		fileName string
		text     string
		error    bool
	}{
		{"testExistingFile",
			[]Student{
				{Name: "Примеров Пример Примерович", Mark: 5},
				{Name: "Примеров Пример Примерович", Mark: 6},
				{Name: "Примеров Пример", Mark: 7},
			}, "../../data/my_output_1.tsv",
			"18\t6\t7\nПримеров Пример Примерович\t11\t5.50\nПримеров Пример\t7\t7", false,
		},

		{"testEmptyPath",
			[]Student{
				{Name: "Примеров Пример Примерович", Mark: 5},
				{Name: "Примеров Пример Примерович", Mark: 6},
				{Name: "Примеров Пример", Mark: 7},
			}, "",
			"", true,
		},
	}
	for _, tc := range testCasesStudents {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()

			writer, _ := os.OpenFile(tc.fileName, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
			defer func(writer *os.File) {
				err := writer.Close()
				if err != nil {
					require.True(t, tc.error)
				}
			}(writer)

			students := Students{tc.students}

			err := WriteStudentsStatistic(writer, &students)
			if err != nil {
				require.True(t, tc.error)
				return
			}

			content, err := os.ReadFile(tc.fileName)
			if err != nil {
				require.True(t, tc.error)
				return
			}

			require.Equal(t, tc.text, string(content))
		})
	}
}
