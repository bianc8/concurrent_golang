# Programmazione Concorrente in GoLang

Ci sono 4 esercizi da consegnare:



# Consegna 1

Scrivete un programma che simuli una agenzia di viaggi che deve gestire le prenotazioni per due
diversi viaggi da parte di 7 clienti. Ogni cliente fa una prenotazione per un viaggio in una delle due
mete disponibili (Spagna e Francia), ognuna delle quali ha un numero minimo di partecipanti per
essere confermata (rispettivamente 4 e 2).

● Creare la struttura Cliente col relativo campo “nome”.

● Creare la struttura Viaggio col rispettivo campo “meta”.

● Creare la function prenota, che prende come input una persona e che prenota uno a caso dei due
viaggi.

● Creare una function stampaPartecipanti che alla fine del processo stampa quali viaggi sono
confermati e quali persone vanno dove.

● Ogni persona può prenotarsi al viaggio contemporaneamente.

● Create tutte le classi e function che vi servono, ma mantenete la struttura data dalle due strutture e
le due function che ho elencato sopra.



# Consegna 2

Scrivete un programma che simuli l’ordinazione, la cottura e l’uscita dei piatti in un ristorante. 10 clienti
ordinano contemporaneamente i loro piatti. In cucina vengono preparati in un massimo di 3 alla volta,
essendoci solo 3 fornelli. Il tempo necessario per preparare ogni piatto è fra i 4 e i 6 secondi. Dopo che
un piatto viene preparato, viene portato fuori da un cameriere, che impiega 3 secondi a portarlo fuori. Ci
sono solamente 2 camerieri nel ristorante.

● Creare la strutture Piatto e Cameriere col relativo campo “nome”.

● Creare le funzioni ordina che aggiunge il piatto a un buffer di piatti da fare; creare la function cucina che
cucina ogni piatto e lo mette in lista per essere consegnato; creare la function consegna che fa uscire
un piatto dalla cucina.

● Ogni cameriere può portare solo un piatto alla volta.

● Usate buffered channels per svolgere il compito.

● Attenzione: se per cucinare un piatto lo mandate nel buffer fornello di capienza 3 e lo ritirate dopo 3
secondi, non è detto che ritiriate lo stesso piatto che avete messo sul fornello. Tenetelo in memoria.



# Consegna 3

Scrivete un programma che simuli un lavoro fatto da tre operai, ognuno dei quali deve usare un
martello, un cacciavite e un trapano per fare un lavoro. Devono usare il cacciavite DOPO il trapano e
il martello in un qualsiasi momento. Ovviamente, possono fare solo un lavoro alla volta. Gli attrezzi a
disposizione sono: due trapani, un martello e un cacciavite, quindi I tre operai devono aspettare di
avere a disposizione gli attrezzi per usarli. Modellate questa situazione minimizzando il più possibile le
attese.

● Creare la struttura Operaio col relativo campo “nome”.

● Creare la strutture Martello, Cacciavite e Trapano che devono essere “prese” dagli operai.

● Nelle function che creerete, inserite una stampa che dica quando l’operaio x ha preso l’oggetto y e
quando ha finito di usarlo.

● Hint sulla logica: ogni operaio può avere solo un oggetto alla volta e ogni oggetto può essere in mano
a un solo operaio.

● Per assicurarmi che ogni operaio abbia in mano un solo oggetto, posso mettere ogni operaio in un
channel, e prima di cercare di prendere un oggetto...



# Consegna 4

Vi è dato un programma (che trovate su Moodle: tunnelBug.go) che simuli la seguente situazione:

● Ci sono due gruppi di palline G1 e G2 in due luoghi diversi L1 e L2 uniti da un tunnel. In L1 e in
L2 ci sono due persone P1 e P2.

● La persona P1 vuole lanciare tutte le palline in G1 da L1 a L2, e viceversa P2 vuole lanciare
lanciare le palline in G2 da L2 a L1.

● Il tunnel è stretto, ci può passare solo una pallina alla volta. Se due palline vengono lanciate nel
tunnel contemporaneamente, tornano al punto di partenza (immediatamente). Una pallina
attraversa il tunnel in un secondo (time.Sleep(time.Second)).

● Una persona non può lanciare una pallina finché quella che ha lanciato precedentemente non è
arrivata a destinazione o non ha incontrato una pallina che andava in senso contrario.

● Ci sono due gruppi di palline e due routine che lanciano le palline da un capo all’altro. Le routine
attendono un tempo casuale (time.Sleep(time.Duration(rand.Intn(2))*time.Second)) prima di
lanciare una nuova palla. Le routine finiscono quando nel relativo gruppo le palline finiscono.


## Cosa Fare?

Debuggate i deadlock! Ci sono errori nel codice. Fate esperimenti, vedete cosa succede,
ripercorrete la logica e rendete mutualmente esclusive la parti di codice che devono avvenire
una alla volta. Possono esserci soluzioni diverse e non mi aspetto che raggiungiate tutti la
stessa.

## Hint

Nel codice, se il tunnel è libero, aspetto un oggetto che mi viene inviato se l’altra routine trova il
tunnel occupato. E se l’altra routine trova il tunnel occupato ma mi invia il messaggio quando
l’altra routine è già andata oltre? Una modifica semplice evita un deadlock, ma crea altri
problemi.

● Una possibile soluzione prevede tre modifiche principali al codice: un altro channel per essere
sicuro che solo un thread stia lavorando su una certa parte di codice; la modifica di un channel
esistente e una select.

● Siate creativi, non è detto che siano le uniche possibiltà. Provate a decommentare gli sleep per
vedere se trovate deadlock sistematici e provate a aggiungere sleep prima e dopo le aprti di
codice che inserite per vedere se trovate deadlock sistematici.