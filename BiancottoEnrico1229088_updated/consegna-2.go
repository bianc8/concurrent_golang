/*
	@author Enrico Biancotto 1229088

	Consegna 2

	Scrivete un programma che simuli l’ordinazione, la cottura e l’uscita dei piatti in un ristorante.
	10 clienti ordinano contemporaneamente i loro piatti.
	In cucina vengono preparati in un massimo di 3 alla volta, essendoci solo 3 fornelli.
	Il tempo necessario per preparare ogni piatto è fra i 4 e i 6 secondi.
	Dopo che un piatto viene preparato, viene portato fuori da un cameriere,
	che impiega 3 secondi a portarlo fuori.
	Ci sono solamente 2 camerieri nel ristorante.

	● Creare la strutture Piatto e Cameriere col relativo campo “nome”.

	● Creare la funziona ordina che aggiunge il piatto a un buffer di piatti da fare;

	● creare la function cucina che cucina ogni piatto e lo mette in lista per essere consegnato;

	● creare la function consegna che fa uscire un piatto dalla cucina.

	● Ogni cameriere può portare solo un piatto alla volta.

	● Usate buffered channels per svolgere il compito.
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Cameriere struct {
	nome string // nome del Cameriere
}
type Piatto struct {
	nome string // nome del Piatto
}

// var globale WaitGroup
var wg sync.WaitGroup

// canali globali che garantiscono mutex
var daFare chan Piatto // canale piatti da cucinare
var daConsegnare chan Piatto // canale piatti cucinati, da consegnare
var camerieri chan Cameriere // canale con camerieri del ristorante

func ordina(p Piatto) {
	defer wg.Done()
	daFare <- p
	fmt.Println(p.nome, "ORDINATO")
}

func cucina() {
	defer wg.Done()
	fornelli := make(chan int, 3)

	for i:= 0; i < 10; i++ { // aspetto 10 clienti
		p := <- daFare
		fornelli <- 1 // anziche mettere un piatto uso un intero simbolico

		start := time.Now()
		time.Sleep(time.Duration(rand.Intn(7 - 4) + 4) * time.Second)

		<- fornelli // svuoto fornello da "piatto" simbolico
		daConsegnare <- p
		fmt.Println(p.nome, "CUCINATO in ",time.Since(start), "s")
	}
}

func uscita() {
	defer wg.Done()
	
	for i:= 0; i < 10; i++ { // aspetto 10 clienti
		tmp := <- camerieri // garantisce mutex su camerieri
		p := <- daConsegnare

		time.Sleep(3 * time.Second)
		
		camerieri <- tmp
		fmt.Println(p.nome, " CONSEGNATO da ",tmp.nome)
	}
}

func main() {
	// init channels
	daFare = make(chan Piatto)
	daConsegnare = make(chan Piatto)
	camerieri = make(chan Cameriere, 2)	// ci sono solo due camerieri nel ristorante

	// init camerieri
	for i:=1; i < 3; i++ {
		nome := fmt.Sprint("cameriere ", i)
		camerieri <- Cameriere{nome}
	}
	
	// init nomi piatti (extra)
	var piatti []Piatto
	for i := 1; i <= 10; i++ {
		nome := fmt.Sprint("piatto ", i)
		piatti = append(piatti, Piatto{nome})
	}


	// 10 ordini
	for i:=0; i < 10; i++ {
		wg.Add(1)
		go ordina(piatti[i])
	}

	wg.Add(2)
	go cucina()
	go uscita()

	wg.Wait()
	close(daFare)
	close(daConsegnare)
	close(camerieri)
}