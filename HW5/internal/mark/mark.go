package mark

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Student struct {
	Name string
	Mark int
}

type StudentsStatistic interface {
	SummaryByStudent(student string) (int, bool)     // default_value, false - если студента нет
	AverageByStudent(student string) (float32, bool) // default_value, false - если студента нет
	Students() []string
	Summary() int
	Median() int
	MostFrequent() int
}

type Students struct {
	studentsSlice []Student
}

func (students *Students) SummaryByStudent(student string) (int, bool) {
	var sum int
	ans := false
	for _, locStudent := range students.studentsSlice {
		if locStudent.Name == student {
			sum += locStudent.Mark
			ans = true
		}
	}
	return sum, ans
}

func (students *Students) AverageByStudent(student string) (float32, bool) {
	var sum int
	var count int
	for _, locStudent := range students.studentsSlice {
		if locStudent.Name == student {
			sum += locStudent.Mark
			count++
		}
	}
	var average float32
	if count == 0 {
		return average, false
	} else {
		return float32(math.Round(float64(sum)/float64(count)*100)) / 100, true
	}
}

func (students *Students) Students() []string {
	ans := map[string]int{}
	for _, locStudent := range students.studentsSlice {
		ans[locStudent.Name] += locStudent.Mark
	}
	keys := make([]string, len(ans))

	i := 0
	for key := range ans {
		keys[i] = key
		i++
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return ans[keys[i]] > ans[keys[j]]
	})

	return keys
}

func (students *Students) Summary() int {
	var sum int
	for _, locStudent := range students.studentsSlice {
		sum += locStudent.Mark
	}
	return sum
}

func (students *Students) Median() int {
	length := len(students.studentsSlice)
	if length == 0 {
		return 0
	}
	keys := make([]int, length)
	for i, locStudent := range students.studentsSlice {
		keys[i] = locStudent.Mark
	}
	sort.Ints(keys)
	return keys[length/2]
}

func (students *Students) MostFrequent() int {
	ansCount := 0
	var retValue int
	ans := map[int]int{}

	for _, locStudent := range students.studentsSlice {
		ans[locStudent.Mark] += 1
		if ans[locStudent.Mark] > ansCount || ans[locStudent.Mark] == ansCount && locStudent.Mark > retValue {
			ansCount, retValue = ans[locStudent.Mark], locStudent.Mark
		}
	}

	return retValue
}

func ReadStudentsStatistic(reader io.Reader) (StudentsStatistic, error) {
	var students []Student
	bufReader := bufio.NewReader(reader)
	for {
		name, err := bufReader.ReadString('\t')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		for strings.Contains(name, "\n") {
			name = name[strings.Index(name, "\n")+1:]
		}
		var a int
		mark, err := bufReader.ReadString('\n')
		if err != nil {
			var err2 error
			a, err2 = strconv.Atoi(mark)
			if err2 != nil {
				return nil, err
			}
		} else {
			var err2 error
			a, err2 = strconv.Atoi(mark[:len(mark)-1])
			if err2 != nil {
				return nil, err
			}
		}
		if a <= 10 && a >= 1 {
			students = append(students, Student{name[:len(name)-1], a})
		}
	}
	return &Students{students}, nil
}

func WriteStudentsStatistic(writer io.Writer, statistic StudentsStatistic) error {
	summary, median, frequent := statistic.Summary(), statistic.Median(), statistic.MostFrequent()
	_, printErr1 := fmt.Fprintf(writer, "%d\t%d\t%d\n", summary, median, frequent)
	if printErr1 != nil {
		return printErr1
	}

	myStrings := make([]string, 0, len(statistic.Students()))
	for _, student := range statistic.Students() {
		summaryByStudent, _ := statistic.SummaryByStudent(student)
		average, _ := statistic.AverageByStudent(student)
		if average == (float32(int(average))) {
			myStrings = append(myStrings, fmt.Sprintf("%s\t%d\t%g", student, summaryByStudent, average))
		} else {
			myStrings = append(myStrings, fmt.Sprintf("%s\t%d\t%.2f", student, summaryByStudent, average))
		}
	}
	_, printErr2 := fmt.Fprint(writer, strings.Join(myStrings, "\n"))
	if printErr2 != nil {
		return printErr2
	}
	return nil
}
