# Oppgave 1
For å sjekke hvor mange prosesser som kjører på vår datamaskin trykker vi på "ctrl + alt + delete" og åpner oppgavebehandleren. Deretter trykker vi på "ytelse" i navigasjonsbaren. Når vi har gjort dette kan vi se hvor mange prosesser som kjører i vårt program, sammen med en del annen informasjon. For å sjekke hvor mange prosesser som kjører i den virtuelle serveren bruker vi kommandoen "top" i PuTTY.

# Oppgave 2
1. Først må vi installere Golang på den virtuelle serveren ved bruk av kommandoen "sudo apt install golang-go".
2. For å åpne teksteditoren i ubuntu serveren bruker vi kommandoen "sudo nano".
3. Vi skriver deretter inn koden som er vist under, og lagrer ved å trykke "ctrl + X" etterfulgt av "Y" for å gå videre til å kunne skrive inn filnavn og trykke enter for å lagre.
<br>![Kode](http://i.imgur.com/C6cL7xr.png)
4. Vi kan deretter teste om filen fungerer ved å bruke kommandoen "go run hello.go"
5. For å gjøre hello.go om til en fil som er kjørbar i Windows bruker man kommdandoen "GOOS=windows GOARCH=386 go build -o hello.exe hello.go"
6. For å laste ned filen logger vi inn på serveren via et program som heter WinSCP. Ved bruk av dette programmet kan vi enkelt laste ned filen hello.exe ved å dra filen til vårt skrivebord eller ønsket destinasjon.
7. for å kjøre filen i windows åpner vi cmd og lokaliserer filen hello.exe, og kjører den ved å skrive "hello.exe".

# Oppgave 3
Del 1 (main_sum.go):
1. For å kjøre programmet må vi lokalisere main_sum.go.
2. Deretter skriver vi følgende i kommandovinduet "go run main_sum.go arg1 arg2". Argumentene som fylles inn må være integer (float64), som tillater tall opp til 1.7976931348623157e+308. 
3. Dette gir oss summen av de to argumentene. 

Del 2 (sum_test.go):
1. Lokaliser sum_test.go
2. For å kjøre testen skriver du "go test" i kommandovinduet. Dette returnerer verdier som overløper de forskjellige typene. 

# Oppgave 4
Del 1 (main.go):
1. Lokaliser main.go i mappen "Oppgave 4".
2. for å kjøre filen, skriv følgende i kommandovinduet "go run main.go". Dette vil sortere en hardkodet tallrekke og printe ut den predefinerte tallrekken og den sorterte tallrekken. 

Del 2(Algorithms):
1. Lokaliser filen "sorting_test.go" i mappen Algorithms.
2. Skriv i kommandovinduet "go test-bench=.". Dette vil kjøre benchmarktesten vår.
