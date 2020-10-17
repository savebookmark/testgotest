package mutexlock

import (
	"fmt"
	"math/rand"
	"sync"
)

type Account struct {
	Balance int
	Mutex   *sync.Mutex
}

func (a *Account) Widthdraw(val int) {
	a.Mutex.Lock()
	a.Balance -= val
	a.Mutex.Unlock()
}

func (a *Account) Deposit(val int) {
	a.Mutex.Lock()
	a.Balance += val
	a.Mutex.Unlock()
}

func (a *Account) Balancefunc() int {
	a.Mutex.Lock()
	balance := a.Balance
	a.Mutex.Unlock()
	return balance
}

var Accounts []*Account
var GlobalLock *sync.Mutex

func Transfer(sender, receiver int, money int) {
	GlobalLock.Lock()
	Accounts[sender].Widthdraw(money)
	Accounts[receiver].Deposit(money)
	GlobalLock.Unlock()
}

func GetTotalBalance() int {
	GlobalLock.Lock()
	total := 0
	for i := 0; i < len(Accounts); i++ {
		total += Accounts[i].Balancefunc()
	}
	GlobalLock.Unlock()
	return total
}

func RandomTranfer() {
	var sender, balance int
	for {
		sender = rand.Intn(len(Accounts))
		balance = Accounts[sender].Balancefunc()
		if balance > 0 {
			break
		}
	}

	var receiver int
	for {
		receiver = rand.Intn(len(Accounts))
		if sender != receiver {
			break
		}
	}

	money := rand.Intn(balance)
	Transfer(sender, receiver, money)
}

func GoTransfer() {
	for {
		RandomTranfer()
	}
}

func PrintTotalBalance() {
	fmt.Printf("Total: %d\n", GetTotalBalance())
}
