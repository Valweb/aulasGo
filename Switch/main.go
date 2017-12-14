package main

import "fmt"

func main() {
	numero := 3
	fmt.Print("O numero se esceve", numero, "se escreve assim")
	switch numero {
	case 1:
		fmt.Println("um.")
	case 2:
		fmt.Println("dois.")
	case 3:
		fmt.Println("tres.")
	case 4:
		fmt.Println("quatro.")
	case 5:
		fmt.Println("cinco.")

	}
}
