package pointers

type Bitcoin int

type Wallet struct {
	// if a symbol (variables, types, functions et al) starts with a lowercase symbol then it is private 
	// outside the package it's defined in.
	balance Bitcoin
}

// take a ponter to the Wallet so that we can change the original values within it.
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	// no need for `(*w).balance`
	// struct pointers are automatically dereferenced
	return w.balance
}