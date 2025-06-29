package main

import "fmt"

type BankAccount struct {
	Owner   string
	Balance float64
}

func (a *BankAccount) Deposit(amount float64) bool {
	if amount > 0 {
		a.Balance += amount
		fmt.Printf("Пополнение на %.2f. Новый баланс: %.2f\n", amount, a.Balance)
		return true
	}
	fmt.Println("Ошибка: сумма пополнения должна быть положительной")
	return false
}

func (a *BankAccount) Withdraw(amount float64) bool {
	if amount > 0 && amount <= a.Balance {
		a.Balance -= amount
		fmt.Printf("Снятие %.2f. Новый баланс: %.2f\n", amount, a.Balance)
		return true
	}
	fmt.Println("Ошибка: недостаточно средств или неверная сумма")
	return false
}

func (a BankAccount) CheckBalance() float64 {
	fmt.Printf("Текущий баланс %s: %.2f\n", a.Owner, a.Balance)
	return a.Balance
}

func main() {
	account := BankAccount{Owner: "Миша", Balance: 1263}
	account.CheckBalance()
	account.Deposit(3000)
	account.Withdraw(5000)
	account.Withdraw(4000)
	account.CheckBalance()
}
