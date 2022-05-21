/*
	@author Enrico Biancotto 1229088

	Consegna 3

	Scrivete un programma che simuli un lavoro fatto da tre operai,
	ognuno dei quali deve usare un martello, un cacciavite e un trapano per fare un lavoro.
	Devono usare il cacciavite DOPO il trapano e il martello in un qualsiasi momento.
	Ovviamente, possono fare solo un lavoro alla volta.
	Gli attrezzi a disposizione sono: due trapani, un martello e un cacciavite.
	Quindi i tre operai devono aspettare di avere a disposizione gli attrezzi per usarli.

	Modellate questa situazione minimizzando il più possibile le attese.

	● Creare la struttura Operaio col relativo campo “nome”.

	● Creare la strutture Martello, Cacciavite e Trapano che devono essere “prese” dagli operai.

	● Nelle function che creerete, inserite una stampa che dica quando l’operaio x ha preso l’oggetto y e quando ha finito di usarlo.
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// l'operaio non può usare il cacciavite finchè sia o.martello che o.trapano sono entrambe true
type Operaio struct {
	nome string		// nome dell'operaio
	martello bool	// operaio.martello == true una volta che l'operaio ha usato il martello
	trapano bool	// operaio.trapano == true una volta che l'operaio ha usato il trapano
}

type Martello struct {
	nome string	// identificativo del Martello
}
type Cacciavite struct {
	nome string	// identificativo del Cacciavite
}
type Trapano struct {
	nome string // identificativo del Trapano
}

var wg sync.WaitGroup

var operai chan Operaio
var martelli chan Martello // canale che conterra' 1 martello
var trapani chan Trapano 	// canale che conterra' 2 trapani
var cacciaviti chan Cacciavite // canale che conterrà 1 cacciavite

func main() {
	// riempio canale operai con 3 Operai 
	operai = make(chan Operaio, 3)
	for i:=1; i<4; i++ {
		operai <- Operaio{nome: fmt.Sprint("Operaio ",i), martello: false, trapano: false}
	}

	// riempie canale martelli con 1 martello
	martelli = make(chan Martello, 1)
	martelli <- Martello{"Martello 1"}

	// riempie canale trapani con 2 trapani
	trapani = make(chan Trapano, 2)
	for i:= 1; i < 3; i++ {
		trapani <- Trapano{fmt.Sprint("Trapano ", i)}
	}
	
	// riempie canale cacciaviti con 1 cacciavite
	cacciaviti = make(chan Cacciavite, 1)
	cacciaviti <- Cacciavite{"Cacciavite 1"}

	// lancia X lavori
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go lavora()
		time.Sleep(200* time.Microsecond)
	}

	wg.Wait()
	close(operai)
	close(martelli)
	close(trapani)
	close(cacciaviti)
}

// simula lavoro di un operaio: un martello, un trapano, un cacciavite
func lavora() {
	defer wg.Done()
	wg.Add(3)
	go usaMartello()
	go usaTrapano()
	go usaCacciavite()
}


func usaCacciavite() {
	defer wg.Done()
	done := false
	for done == false {
		select {
		case ope := <-operai:
			// mutex garantita perchè operaio è stato tolto da channel (sync auto da Go)
			// questo if garantisce che l'operaio usi il cacciavite solo dopo aver usato il martello e il trapano,
			// questo approcio garantisce attese minime sugli atrezzi,
			//  quindi un operaio non blocca altri se non ha l'atrezzo che gli serve
			// ma un operaio può aspettare molti cicli prima di finire un lavoro(1 martello + 1 trapano + 1 cacciavite)
			if (ope.martello == true) && (ope.trapano == true) {
				select {
				case atrezzo := <- cacciaviti:	// garantisce mutex sull'unico cacciavite
					done = true
					rand.Seed(time.Now().UnixNano())
					start := time.Now()
					time.Sleep(time.Duration(rand.Intn(2) + 1) * time.Second)
					
					// mutex su variabili garantita da channel operai
					ope.martello = false
					ope.trapano = false
					
					fmt.Println(ope.nome, "ha usato", atrezzo.nome, " per ",time.Since(start),"s")
					
					cacciaviti <- atrezzo
				}
			}
			operai <- ope
		}
	} 
}

// simile a usaCacciavite
func usaMartello() {
	defer wg.Done()
	done := false
	for done == false {
		select {
		case ope := <-operai:
			// mutex garantita perchè operaio è stato tolto da channel (sync auto da Go)
			if ope.martello == false {
				select {
				case atrezzo := <-martelli:	// garantisce mutex sull'unico martello
					done = true
					rand.Seed(time.Now().UnixNano())
					start := time.Now()
					time.Sleep(time.Duration(rand.Intn(2) + 1) * time.Second)
					
					// mutex su variabili garantita da channel operai
					ope.martello = true
					
					fmt.Println(ope.nome, "ha usato", atrezzo.nome, " per ",time.Since(start),"s")
					
					martelli <- atrezzo
				}
			}
			operai <- ope
		}
	}
}

// identica a usaMartello, cambia solo il canale In/Out per l'atrezzo
func usaTrapano() {
	defer wg.Done()
	done := false
	for done == false {
		select {
		case ope := <-operai:
			// mutex garantita perchè operaio è stato tolto da channel (sync auto da Go)
			if ope.trapano == false {
				select {
				case atrezzo := <-trapani:   // garantisce mutex sui trapani
					done = true
					rand.Seed(time.Now().UnixNano())
					start := time.Now()
					time.Sleep(time.Duration(rand.Intn(2) + 1) * time.Second)
					
					// mutex su variabili garantita da channel operai
					ope.trapano = true
				
					fmt.Println(ope.nome, "ha usato", atrezzo.nome, " per ",time.Since(start),"s")

					trapani <- atrezzo
				}
			}
			operai <- ope
		}
	}
}
