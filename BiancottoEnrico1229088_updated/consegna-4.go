/*
	@author Enrico Biancotto 1229088

	Consegna 4

	Vi è dato un programma (che trovate su Moodle: tunnelBug.go) che simuli la seguente situazione:

	● Ci sono due gruppi di palline G1 e G2 in due luoghi diversi L1 e L2 uniti da un tunnel.
	In L1 e in L2 ci sono due persone P1 e P2.

	● La persona P1 vuole lanciare tutte le palline in G1 da L1 a L2, e viceversa P2 vuole lanciare lanciare le palline in G2 da L2 a L1.

	● Il tunnel è stretto, ci può passare solo una pallina alla volta.
	Se due palline vengono lanciate nel tunnel contemporaneamente, tornano al punto di partenza (immediatamente).
	Una pallina attraversa il tunnel in un secondo (time.Sleep(time.Second)).

	● Una persona non può lanciare una pallina finché quella che ha lanciato precedentemente non è arrivata a destinazione o non ha incontrato una pallina che andava in senso contrario.

	● Ci sono due gruppi di palline e due routine che lanciano le palline da un capo all’altro.
	Le routine attendono un tempo casuale prima di lanciare una nuova palla.

	Le routine finiscono quando nel relativo gruppo le palline finiscono.
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// var globale WaitGroup
var wg sync.WaitGroup

type Gruppo struct {
    nome string
    nPalline int
}

type Tunnel struct {
    libero bool
}


func transumanza(g Gruppo, t chan bool){
	defer wg.Done()
	for g.nPalline > 0 {
		time.Sleep(time.Duration(rand.Intn(2))*time.Second)
		mandaPersona(&g, t)
	}
}

// sincronizzazione garantita dal buffered channel
func mandaPersona(g *Gruppo, t chan bool){
	select{
	case t <- true:
		time.Sleep(time.Second)
		g.nPalline--
		fmt.Println("\nNe rimangono ", g.nPalline, " in ", g.nome)
		<- t
	default:
		fmt.Println("\nScontro del ",g.nome," ")
		time.Sleep(time.Duration(1)*time.Second)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	gruppo1 := Gruppo{"Gruppo 1", 5}
	gruppo2 := Gruppo{"Gruppo 2", 5}
	
	// fulcro della sincronizzazione
	tunnel := make(chan bool, 1)
	
	wg.Add(2)
	go transumanza(gruppo1, tunnel)
	go transumanza(gruppo2, tunnel)
	
	wg.Wait()
	close(tunnel)
}