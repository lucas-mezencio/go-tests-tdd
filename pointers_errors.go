package main

import "errors"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

var ErrInsufficientFunds = errors.New("cannot withdraw, unsufficient funds")

// Deposit sem passar o ponteiro de wallet ele chama uma cópia
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	// (*w).balance poderia ser usado
	// mas struct pointers sao automaticamente desreferenciados
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
