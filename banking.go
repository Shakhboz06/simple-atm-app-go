package main

import (
	"functions.go/fileops"
	"fmt"
)

func main() {
	// accountBalance := 3000.89
	accountBalance, err := fileops.ReadFile("accountBalance.txt")
	if err != nil {
		fmt.Println("ERROR", err)
		panic("Cannot be continued! ERRORRRR")
	}
	
	fmt.Println("Welcome to Go Banking!")

	for {
		ChooseOptions()

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
				fmt.Printf("in your balance left: %.2f\n", accountBalance)
				fileops.File(accountBalance, "accountBalance.txt")
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
			fmt.Printf("In your balance there is %.2f\n", accountBalance)
			fileops.File(accountBalance,"accountBalance.txt")
		} else if choice == 4 {
			fmt.Println("You exited!")
			break
		} else {
			fmt.Println("EEEE wrong. Try again!")
		}

	}
}
