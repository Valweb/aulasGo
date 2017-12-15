package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	numero := 3
	fmt.Print("O numero  ", numero, "  se escreve assim =>")
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
	fmt.Println()
	fmt.Println("Você é da família do Unix?")

	//Verifica o SOS
	switch runtime.GOOS {

	case "darwin":
		fallthrough // Segue para o próximo...
	case "freebsd":
		fallthrough
	case "linux":
		fmt.Println("Sim, sou Ubuntu! =D <3")
	default:
		fmt.Println("Não vamos falara sobre isso.")
	}

	fmt.Println()

	//Weekday Verifica o dia da semana
	fmt.Printf("hoje é: %s\n\r", time.Now().Weekday())
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: // o Go aceita mais de uma expressão no case
		fmt.Println("Ok, você pode dormir até mais tarde!")
	default:
		fmt.Println("Vamos lá, temos muito trabalho para executar")
	}

	fmt.Println()

	numero = 10
	fmt.Println("Este número cabe em um dígito?")

	//colocar a condição no case e não no swtich
	switch {

	case numero < 10:
		fmt.Println("Sim")
	case numero >= 10 && numero < 100:
		fmt.Println("Serve dois dígitos?")
	}

	fmt.Println()

	cidade := "BH"
	fmt.Println("Você é Minas Gerais?")

	switch cidade {

	case "RJ":
		fallthrough
	case "SP":
		fallthrough
	case "BH":
		fmt.Println("Sim, sou de Belo Horizonte!")
	default:
		fmt.Println("Sou do mundo!")
	}
}
