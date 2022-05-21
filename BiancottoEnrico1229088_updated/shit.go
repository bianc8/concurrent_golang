package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var viaggi[2] Viaggio
var clienti[7]Cliente
var channel1 = make(chan bool)
var channel2 = make(chan bool)

type Cliente struct {
	nome string
	presenza bool
}

/***
spagna  = 0  / partecipanti = 4
francia = 1  / partecipanti = 2
*/
type Viaggio struct {
	meta string
	letsgo int
	clienti [7]Cliente
}

func stampaPartecipanti(){
	if viaggi[0].letsgo >= 4  {
		for i:=0;i<7;i++ {
			if viaggi[0].clienti[i].presenza == true {
				fmt.Println("cliente ",viaggi[0].clienti[i].nome ,"ha prenotato per la francia")
			}
		}
	}else {
		fmt.Println("non ci sono abbastanza iscritti per la francia")
	}
	if viaggi[1].letsgo >= 2  {
		for i:=0;i<7;i++ {
			if viaggi[0].clienti[i].presenza == true {
				fmt.Println("cliente ",viaggi[1].clienti[i].nome ,"ha prenotato per la spagna")
			}
		}
	}else {
		fmt.Println("non ci sono abbastanza iscritti per la spagna")
	}
}

func prenota(c Cliente) {
	go func() {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(500)+500) * time.Millisecond)
		channel1 <- true
	}()
	go func() {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(500)+500) * time.Millisecond)
		channel1 <- false
	}()
}

func main(){
	viaggi[0].meta = "spagna"
	viaggi[1].meta = "francia"
	clienti[0].nome = "Giovanni"
	clienti[1].nome = "Michele"
	clienti[2].nome = "Giacomo"
	clienti[3].nome = "Samuele"
	clienti[4].nome = "Silvia"
	clienti[5].nome = "Stefania"
	clienti[6].nome = "Daniela"
	viaggi[0].clienti = clienti
	viaggi[1].clienti = clienti

	for i:=0; i<7; i++ {
		go prenota(clienti[i])
		select {
		case message1 := <-channel1:
			viaggi[0].letsgo++
			viaggi[0].clienti[i].presenza = message1
		case message2 := <-channel2:
			viaggi[1].letsgo++
			viaggi[1].clienti[i].presenza = message2
		}
	}
	stampaPartecipanti()
	time.Sleep(time.Second*2)
}