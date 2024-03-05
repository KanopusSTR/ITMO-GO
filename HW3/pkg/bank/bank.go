package bank

import (
	"fmt"
	"strings"
	"time"
)

const (
	TopUpOp OperationType = iota
	WithdrawOp
)

type OperationType int64

type Clock interface {
	Now() time.Time
}

func NewRealTime() *RealClock {
	return &RealClock{}
}

type RealClock struct{}

func (account *RealClock) Now() time.Time {
	return time.Now()
}

type Operation struct {
	OpTime   time.Time
	OpType   OperationType
	OpAmount int
	Balance  int
}

func (o Operation) String() string {
	var format string
	if o.OpType == TopUpOp {
		format = `%s +%d %d`
	} else {
		format = `%s -%d %d`
	}
	return fmt.Sprintf(format, o.OpTime.String()[:19], o.OpAmount, o.Balance)
}

type Account interface {
	TopUp(amount int) bool
	Withdraw(amount int) bool
	Operations() []Operation
	Statement() string
	Balance() int
}

type RealAccount struct {
	clock      Clock
	balance    int
	operations []Operation
}

func (account *RealAccount) TopUp(amount int) bool {
	if amount <= 0 {
		return false
	}
	account.balance += amount
	account.operations = append(account.operations, Operation{
		OpTime:   account.clock.Now(),
		OpType:   TopUpOp,
		OpAmount: amount,
		Balance:  account.balance,
	})
	return true
}

func (account *RealAccount) Withdraw(amount int) bool {
	if amount <= 0 || account.balance < amount {
		return false
	}
	account.balance -= amount
	account.operations = append(account.operations, Operation{
		OpTime:   account.clock.Now(),
		OpType:   WithdrawOp,
		OpAmount: amount,
		Balance:  account.balance,
	})
	return true
}

func (account *RealAccount) Operations() []Operation {
	return account.operations
}

func (account *RealAccount) Statement() string {
	var answer []string
	for _, operation := range account.operations {
		strings.Join(answer, operation.String())
		answer = append(answer, operation.String())
	}
	return strings.Join(answer, "\n")
}

func (account *RealAccount) Balance() int {
	return account.balance
}

func NewAccount(clock Clock) Account {
	return &RealAccount{
		clock:      clock,
		balance:    0,
		operations: []Operation{},
	}
}
