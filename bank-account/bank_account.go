package account

import "sync"

// Account simulates an ordinary bank account
type Account struct {
	//id		int64 // random id must be uniq
	mu      sync.Mutex
	balance int64
	open	bool
}

// Open initilize a bank account with given parameters
func Open (firstBalance int64) *Account {
	if firstBalance < 0 {
		return nil
	}

	return &Account{
		balance: firstBalance,
		open:    true,
	}
}

// Balance gets open account current balance
func (a *Account) Balance() (int64,bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.open == false {
		return 0,false
	}
	return a.balance, true
}

// Close Deactivate the bank account and withdraw the account balance
func (a *Account) Close() (int64,bool){
    a.mu.Lock()
    defer a.mu.Unlock()

    if a.open == false {
    	return 0,false
	}

    payout := a.balance
    a.balance = 0
    a.open = false

	return payout, !a.open
}

// Deposit withdraws or deposits money from an open account. Cant withdraw more than current account balance
func (a *Account) Deposit(amount int64) (int64,bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.open == false {
		return 0,false
	}

	if a.balance+amount < 0 {
		return a.balance,false
	}

	a.balance = a.balance + amount
	return a.balance,true
}