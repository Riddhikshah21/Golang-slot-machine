package main

import "fmt"

func getName() string {
	name := ""
	fmt.Println("Welcome to Riddhi's Casino..")
	fmt.Printf("Enter your name: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		return ""
	}
	fmt.Printf("Welcome %s, let's play!\n", name)
	return name
}

func get_bet(balance uint) uint {
	var bet uint
	for true {
		fmt.Printf("Enter your bet, or 0 to quit (balance: $%d): ", balance)
		fmt.Scan(&bet)
		if bet > uint(balance) {
			fmt.Println("Bet cannot be greater than balance amount!")
		} else {
			break
		}
	}
	return bet
}

func main() {
	balance := uint(200)
	getName()

	for balance > 0 {
		bet := get_bet(balance)
		if bet == 0 {
			break
		}
		balance -= bet
	}
	fmt.Printf("You are left with $%d\n", balance)
}
