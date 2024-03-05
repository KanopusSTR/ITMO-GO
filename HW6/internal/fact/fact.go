package fact

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"sync"
)

type Input struct {
	NumsOfGoroutine int   // n - число горутин
	Numbers         []int // слайс чисел, которые необходимо факторизовать
}

type Factorization interface {
	Work(Input, io.Writer) error
}

type MyFactorization struct {
	mu sync.Mutex
	c  int
}

func SoloFactorization(number int) []string {
	var ans []string

	if number < 0 {
		ans = append(ans, "-1")
		number *= -1
	}

	if number == 1 {
		return append(ans, "1")
	}

	for i := 2; i*i <= number; i++ {
		for number%i == 0 {
			ans = append(ans, fmt.Sprintf("%d", i))
			number = number / i
		}
	}

	if number != 1 {
		ans = append(ans, fmt.Sprintf("%d", number))
	}
	return ans
}

func (myFact *MyFactorization) Work(input Input, writer io.Writer) error {
	var wg sync.WaitGroup

	ch := make(chan int, len(input.Numbers))
	go func() {
		for _, number := range input.Numbers {
			ch <- number
		}
		close(ch)
	}()

	if len(input.Numbers) == 0 {
		return nil
	}
	var mError error

	wg.Add(input.NumsOfGoroutine)
	for i := 0; i < input.NumsOfGoroutine; i++ {
		go func() {
			for number := range ch {
				factorization := SoloFactorization(number)

				var e error

				myFact.mu.Lock()

				myFact.c++

				ans := fmt.Sprintf("line %d, %d = %s",
					myFact.c, number, strings.Join(factorization, " * "))

				if myFact.c != len(input.Numbers) {
					_, e = fmt.Fprintln(writer, ans)
				} else {
					_, e = fmt.Fprint(writer, ans)
				}

				myFact.mu.Unlock()

				mError = errors.Join(mError, e)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	return mError
}

func NewFactorization() *MyFactorization {
	return &MyFactorization{}
}
