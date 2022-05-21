/*
	@author Enrico Biancotto 1229088

	Consegna 1

	Scrivete un programma che simuli una agenzia di viaggi che deve gestire le prenotazioni per due
	diversi viaggi da parte di 7 clienti. Ogni cliente fa una prenotazione per un viaggio in una delle due
	mete disponibili (Spagna e Francia), ognuna delle quali ha un numero minimo di partecipanti per
	essere confermata (rispettivamente 4 e 2).

		● Creare la struttura Cliente col relativo campo “nome”.

		● Creare la struttura Viaggio col rispettivo campo “meta”.

		● Creare la function prenota, che prende come input una persona e che prenota uno a caso dei due viaggi.

		● Creare una function stampaPartecipanti che alla fine del processo stampa quali viaggi sono confermati e quali persone vanno dove.

		● Ogni persona può prenotarsi al viaggio contemporaneamente.

		● Create tutte le classi e function che vi servono, ma mantenete la struttura data dalle due strutture e le due function che ho elencato sopra.
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Cliente struct {
	nome string
}

type Viaggio struct {
	meta string
	minPartecipanti int
	clienti []Cliente
}

var viaggi []Viaggio

// prenota a caso uno dei due viaggi
func prenota(x Cliente) {
	defer wg.Done() // una volta finita una prenotazione avvisa il WaitGroup

	rand.Seed(time.Now().UnixNano())
	var t int = rand.Intn(2)
	
	var tmp []Cliente
	if (t == 1) {
		tmp = viaggi[0].clienti
		tmp = append(tmp, x)
		viaggi[0].clienti = tmp
		fmt.Println("Cliente ",x.nome," ha prenotato in ",viaggi[0].meta)
	} else {
		tmp = viaggi[1].clienti
		tmp = append(tmp, x)
		viaggi[1].clienti = tmp
		fmt.Println("Cliente ",x.nome," ha prenotato in ",viaggi[1].meta)
	}
}

// alla fine del processo stampa quali viaggi sono confermati e quali persone vanno dove
func stampaPartecipanti() {
	for _, elem := range viaggi {
		clienti := elem.clienti
		// clienti deve essere > di minPartecipanti
		if (len(clienti) >= elem.minPartecipanti) {
			fmt.Print("\nViaggio in ",elem.meta," confermato con clienti: ")
			for _, cliente := range clienti {
				fmt.Print(cliente.nome,", ")
			} 
		} else {
			fmt.Println("\nViaggio in ",elem.meta," non confermato")
		}
	}	
	fmt.Println() // output preferences
}

func main() {
	// init clienti
	var clienti []Cliente
	for i:= 1; i < 11; i++ {
		nome := fmt.Sprint("cliente ", i)
		clienti = append(clienti, Cliente{nome})
	}

	// init viaggi
	var tmp []Cliente
	viaggi = append(viaggi, Viaggio{"Spagna",  4, tmp})
	viaggi = append(viaggi, Viaggio{"Francia", 2, tmp})
	
	// lancia go routine prenota
	for i := 0; i < len(clienti); i++ {
		wg.Add(1)
		go prenota(clienti[i])
		time.Sleep(time.Second)
	}

	wg.Wait()

	defer stampaPartecipanti()
}