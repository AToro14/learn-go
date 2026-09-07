package main

import (
	"errors"
)

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

func updateBalance(c *customer, t transaction) error {
	if t.amount < 0 {
		return errors.New("transaction amount cannot be a negative amount")
	}
	if t.transactionType == transactionDeposit {
		c.balance = c.balance + t.amount
		return nil
	} else if t.transactionType == transactionWithdrawal {
		balCheck := c.balance - t.amount
		if balCheck < 0 {
			return errors.New("insufficient funds")
		}
		c.balance = balCheck
		return nil
	} else {
		return errors.New("unknown transaction type")
	}
}
