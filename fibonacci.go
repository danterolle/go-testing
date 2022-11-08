/*

In questa prima sfida verrà scritto un programma per calcolare la sequenza di Fibonacci da un numero.
Questo è un esercizio tipico di scrittura del codice durante l'apprendimento di un nuovo linguaggio.
Si scriverà una funzione che restituisce una slice con tutti i numeri in una sequenza di Fibonacci
risultante dall'operazione di calcolo eseguita su un numero immesso da un utente maggiore di due.

Si supponga che i numeri minori di 2 causino un errore e restituiscano una slice nil.
Tenere presente che la sequenza di Fibonacci è un elenco di numeri in cui ogni numero è
la somma dei due numeri di Fibonacci precedenti.
Ad esempio, la sequenza di numeri per 6 è 1,1,2,3,5,8, per 7 è 1,1,2,3,5,8,13, per 8 è 1,1,2,3,5,8,13,21 e così via.

*/

package main

import (
	"fmt"
)

func fibonacci(number int) []int {
	f := make([]int, number)
	fmt.Println(f)

	if number <= 2 {
		return make([]int, 0)
	}

	f[0], f[1] = 1, 1
	fmt.Println(f)

	for i := 2; i < number; i++ {
		fmt.Println(f[i-1], f[i-2])
		f[i] = f[i-1] + f[i-2]
	}
	return f
}

func main() {
	fmt.Println(fibonacci(6))
}
