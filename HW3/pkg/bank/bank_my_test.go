package bank

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStartBalance(t *testing.T) {
	account := NewAccount(NewMockTime())
	want := 0
	got := account.Balance()
	require.Equal(t, want, got)

}

func TestTopUpNormal(t *testing.T) {
	account := NewAccount(NewMockTime())
	require.True(t, account.TopUp(100))
	require.Equal(t, 100, account.Balance())
}

func TestTopUpError(t *testing.T) {
	account := NewAccount(NewMockTime())
	require.False(t, account.TopUp(-100))
	require.Equal(t, 0, account.Balance())
}

func TestTopUpMany(t *testing.T) {
	account := NewAccount(NewMockTime())
	account.TopUp(100)
	account.TopUp(200)
	require.Equal(t, 300, account.Balance())
}

func TestWithdrawNormal(t *testing.T) {
	account := NewAccount(NewMockTime())
	account.TopUp(1000)
	require.True(t, account.Withdraw(100))
	require.Equal(t, 900, account.Balance())
}

func TestWithdrawError(t *testing.T) {
	account := NewAccount(NewMockTime())
	require.False(t, account.TopUp(-100))
	require.Equal(t, 0, account.Balance())
}

func TestWithdrawMany(t *testing.T) {
	account := NewAccount(NewMockTime())
	account.TopUp(1000)
	account.Withdraw(300)
	account.Withdraw(500)
	require.Equal(t, 200, account.Balance())
}

func TestOperationsNormal(t *testing.T) {
	clock := NewMockTime()
	account := NewAccount(clock)
	account.TopUp(1000)
	require.Equal(t, 1, len(account.Operations()))
	operation := Operation{
		OpTime:   account.Operations()[0].OpTime,
		OpAmount: 1000,
		OpType:   TopUpOp,
		Balance:  1000,
	}
	require.Equal(t, operation, account.Operations()[0])
}

func TestOperationsError(t *testing.T) {
	clock := NewMockTime()
	account := NewAccount(clock)
	account.Withdraw(1000)
	require.Equal(t, 0, len(account.Operations()))
	require.Equal(t, []Operation{}, account.Operations())
}

func TestOperationsMany(t *testing.T) {
	clock := NewMockTime()
	account := NewAccount(clock)
	account.TopUp(1000)
	account.Withdraw(500)
	require.Equal(t, 2, len(account.Operations()))
	operations := []Operation{
		{OpTime: account.Operations()[0].OpTime,
			OpAmount: 1000,
			OpType:   TopUpOp,
			Balance:  1000},
		{OpTime: account.Operations()[1].OpTime,
			OpAmount: 500,
			OpType:   WithdrawOp,
			Balance:  500},
	}
	require.Equal(t, operations, account.Operations())
}

func TestStatementNormal(t *testing.T) {
	clock := NewMockTime()
	account := NewAccount(clock)
	account.TopUp(1000)
	require.Equal(t, "2023-03-18 12:34:07 +1000 1000", account.Statement())
}

func TestStatementError(t *testing.T) {
	clock := NewMockTime()
	account := NewAccount(clock)
	account.Withdraw(1000)
	require.Equal(t, "", account.Statement())
}

func TestStatementMany(t *testing.T) {
	clock := NewMockTime()
	account := NewAccount(clock)
	account.TopUp(1000)
	account.Withdraw(500)
	require.Equal(t, "2023-03-18 12:34:07 +1000 1000\n2023-03-18 12:34:37 -500 500", account.Statement())
}
