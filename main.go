// default pin is 1234

package main

import (
	"fmt"
	"os"
	"strconv"
)

var defaultpin = "1234"
var accessPin string
var availableBal float32

func main() {
	welcome()
	login()
}

func newline(num int) {
	for i := 0; i < num; i++ {
		fmt.Println("\n")
	}
}

func welcome() {
	fmt.Println("*** Welcome to my ATM application ***")
}

func anotherTransaction() {
	var exitInput string

	newline(1)

	fmt.Println("Do you want to perform another transactin? Enter y/n:")

	_, err2 := fmt.Scan(&exitInput)
	if err2 != nil {
		fmt.Println("Invalid Input!")
	}

	if exitInput == "y" {
		menu()
	} else if exitInput == "n" {
		exit()
	} else {
		fmt.Println("Invalid input")
		anotherTransaction()
	}
}

func login() {

	newline(1)
	fmt.Println("Please enter your four digit pin:")

	_, err := fmt.Scan(&accessPin)
	if err != nil {
		fmt.Println("Invalid Pin")
	}

	if accessPin == "1234" {
		newline(1)
		fmt.Println("Change the default pin '1234' to a secure four digit pin before you continue")
		menu()
	} else if accessPin == "0" {
		fmt.Println("Goodbye!")
		os.Exit(0)
	} else {
		fmt.Println("Incorrect pin. Enter correct pin or enter 0 to quit.")
		login()
	}
}

func menu() {
	var input int

	newline(1)

	fmt.Println("Choose a transaction:")
	fmt.Println("1 Change pin\t\t 2 View balance")
	fmt.Println("3 Withdraw cash\t\t 4 Deposit")
	fmt.Println("0 Cancel transaction")

	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Invalid Input")
	}

	switch input {
	case 1:
		changePin()
	case 2:
		balance()
	case 3:
		withdraw()
	case 4:
		deposit()
	case 0:
		exit()
	default:
		fmt.Println("Enter a valid menu number")
		menu()
	}

}

func changePin() {
	var oldPin string

	newline(1)
	fmt.Println("Enter your old pin:")
	_, err := fmt.Scan(&oldPin)
	if err != nil {
		fmt.Println("Invalid pin")
	}

	if oldPin == accessPin {
		fmt.Println("Enter new pin:")
		_, err := fmt.Scan(&accessPin)
		if err != nil {
			fmt.Println("Old pin invalid")
		}
		newline(1)
		fmt.Println("*Pin changed successfully*")
		anotherTransaction()
	} else if oldPin == "0" {
		menu()
	} else {
		newline(1)
		fmt.Println("Old pin incorrect. Try again or enter 0 to return to menu.")
		changePin()
	}
}

func balance() {
	newline(1)
	fmt.Printf("Your available balance is $%v", availableBal)
	anotherTransaction()
}

func withdraw() {
	var input string
	var confirmPin string

	newline(1)
	fmt.Println("Please enter withdrawal amount:")

	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Error: invalid amount")
	}

	// confirming pin before withdrawal
	newline(1)
	fmt.Println("Confirm your pin to proceed:")
	_, err1 := fmt.Scan(&confirmPin)
	if err1 != nil {
		fmt.Println("Error: Invalid input")
	}

	withdrawInput, _ := strconv.ParseFloat(input, 32)

	if float32(withdrawInput) <= availableBal && confirmPin == accessPin {
		fmt.Println("Pin accepted")
		newline(1)
		availableBal = availableBal - float32(withdrawInput)
		fmt.Println("** Please take your cash. **")
		anotherTransaction()
	} else if float32(withdrawInput) > availableBal && confirmPin == accessPin {
		fmt.Println("Pin accepted")
		newline(1)
		fmt.Println("* Insufficient Balance *")
		anotherTransaction()
	} else if confirmPin != accessPin {
		fmt.Println("Incorrect Pin")
		withdraw()
	}
}

func deposit() {
	var input int
	newline(1)

	fmt.Println("Choose a deposit method:")
	fmt.Println("1. Debit Card")
	fmt.Println("2. Bank Transfer")
	fmt.Println("3. Crypto(btc/eth)")

	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Invalid input!")
	}
	newline(1)

	switch input {
	case 1:
		fmt.Println("Please enter deposit amount:")
		_, err := fmt.Scan(&availableBal)
		if err != nil {
			fmt.Println("Invalid input")
		}
		fmt.Println("Deposit Successful!")
		anotherTransaction()
	case 2:
		fmt.Println("Please enter deposit amount:")
		_, err := fmt.Scan(&availableBal)
		if err != nil {
			fmt.Println("Invalid input")
		}
		fmt.Println("Deposit Successful!")
		anotherTransaction()
	case 3:
		fmt.Println("Please enter deposit amount:")
		_, err := fmt.Scan(&availableBal)
		if err != nil {
			fmt.Println("Invalid input")
		}
		fmt.Println("Deposit Successful!")
		anotherTransaction()
	}
}

func exit() {
	fmt.Println("** Goodbye! **")
	os.Exit(0)
}
