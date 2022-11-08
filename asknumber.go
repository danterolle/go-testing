/*

Questo programma richiede un numero e lo stampa. Modificare il codice di esempio come segue:

Chiedere continuamente un numero intero. La condizione di uscita per il ciclo deve essere l'immissione di un numero negativo da parte dell'utente.
Arrestare il programma quando l'utente immette un numero negativo. Stampare quindi l'errore di analisi dello stack.
Quando il numero è 0 stampare 0 is neither negative nor positive. Continuare a chiedere un numero.
Quando il numero è positivo, stampare You entered: X (dove X è il numero immesso). Continuare a chiedere un numero.
Per il momento, ignorare la possibilità che l'utente immetta un valore diverso da un numero intero.

*/

package main

import "fmt"

func main() {
	val := 0
	for {
		fmt.Print("Enter number: ")
		fmt.Scanf("%d", &val)
		fmt.Println("You entered: ", val)
		if val < 0 {
			panic(val)
			return
		} else if val == 0 {
			fmt.Println("0 is neither negative nor positive")
			continue
		}
		continue
	}
}
