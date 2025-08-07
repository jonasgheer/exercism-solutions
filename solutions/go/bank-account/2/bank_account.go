package account

import "sync"

type Account struct {
	mu      sync.Mutex
	balance int64
	active  bool
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{balance: amount, active: true}
}

func (a *Account) Balance() (balance int64, ok bool) {
	if a.active == false {
		return 0, false
	}
	return a.balance, true
}

func (a *Account) Deposit(amount int64) (balance int64, ok bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.active == false {
		return a.balance, false
	}
	if a.balance+amount < 0 {
		return a.balance, false
	}
	a.balance += amount
	return a.balance, true
}

func (a *Account) Close() (balance int64, ok bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.active == false {
		return 0, false
	}
	a.active = false
	balance = a.balance
	a.balance = 0
	return balance, true
}
