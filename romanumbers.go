/*

Scrivere un programma che converte un numero romano come MCLX in 1,160.
Usare una mappa per caricare i numeri romani di base che verranno usati per
convertire un carattere stringa in un numero.
Ad esempio, M sarà una chiave nella mappa e il suo valore sarà 1000.
Usare l'elenco di mapping di caratteri stringa seguente:

M => 1000
D => 500
C => 100
L => 50
X => 10
V => 5
I => 1

Se l'input dell'utente include una lettera diversa dall'elenco precedente, stampare un errore.
Si ricordi che in alcuni casi un numero più piccolo precede un numero maggiore,
quindi non è possibile sommare semplicemente i numeri.
Ad esempio, il numero MCM deve essere stampato come 1,900.

*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(conv("LCVI"))
}

func conv(s string) int {
	alphabet := map[string]int{
		"M": 1000,
		"D": 500,
		"C": 100,
		"L": 50,
		"X": 10,
		"V": 5,
		"I": 1,
	}

	sum := alphabet[string(s[len(s)-1])]
	fmt.Println(len(s), sum)

	for i := len(s) - 2; i >= 0; i-- {
		currentInt := alphabet[string(s[i])]
		nextInt := alphabet[string(s[i+1])]
		fmt.Println(sum, currentInt, nextInt)

		/*
			If the current value is greater than or equal to the value of the symbol to the right, add the current symbol’s value to the total
			If the current value is smaller than the value of the symbol to the right, subtract the current symbol’s value from the total.
		*/

		if currentInt >= nextInt {
			sum += currentInt
		} else {
			sum -= currentInt
		}
	}
	return sum
	//int(math.Abs(float64(sum)))
}
