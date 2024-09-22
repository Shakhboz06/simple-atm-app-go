package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func file (balance float64){
	balanceTxt := fmt.Sprint(balance)
	os.WriteFile("accountBalance.txt", []byte(balanceTxt), 0644)
}

func readFile() (float64, error){
	data, err := os.ReadFile("accountBalance.txt")
	if(err != nil) {
		return 10000, errors.New("Failed to read file!")
	}
	balanceTxt := string(data)
	accountBalance, err := strconv.ParseFloat(balanceTxt, 64)
	if(err != nil) {
		return 10000, errors.New("failed to convert!")
	}
	
	return accountBalance, nil
}
func main() {
	// accountBalance := 3000.89
	accountBalance, err := readFile()
	if err != nil {
		fmt.Println("ERROR", err)
		panic("Cannot be continued! ERRORRRR")
	}
	
	fmt.Println("Welcome to Go Banking!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1.Check Balance")
		fmt.Println("2.Withdraw amount")
		fmt.Println("3.Make a deposit")
		fmt.Println("4.Exit")

		var choice int
		var amount float64
		fmt.Print("Choose one of the options: ")
		fmt.Scan(&choice)
		fmt.Println("Your choice is:", choice)

		if choice == 1 {
			fmt.Println("Your balance is: ", accountBalance)
		} else if choice == 2 {
			fmt.Print("How much you want to withdraw: ")
			fmt.Scan(&amount)
			if amount < accountBalance {
				fmt.Println("You have withdrawn: ", amount)
				accountBalance -= amount
				fmt.Printf("in your balance left: %0.2f\n", accountBalance)
				file(accountBalance)
			} else {
				fmt.Println("You have insufficient amount in your account ")
			}
		} else if choice == 3 {
			fmt.Print("How much you want to deposit: ")
			fmt.Scan(&amount)
			if amount <= 0 {
				fmt.Print("Please enter adequate amount of money\n")
				continue
			}
			fmt.Println("You are depositing: ", amount)
			accountBalance += amount
			fmt.Printf("In your balance there is %0.2f\n", accountBalance)
			file(accountBalance)
		} else if choice == 4 {
			fmt.Println("You exited!")
			break
		} else {
			fmt.Println("EEEE wrong. Try again!")
		}

	}
}
