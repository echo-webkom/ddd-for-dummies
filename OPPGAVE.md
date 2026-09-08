# Oppgave

Denne oppgaven krever at du lager et enkelt API for å liste og booke hoteller (veldig likt README).

## Regler

1. Ingen LLM-autocomplete
2. Chat til spørsmål om Go eller andre vertkøy er lov, men ikke for fullføring av oppgaver.

## 0. Hello world

Standarden for Go apper er at `main` pakken og funksjonen ligger i `cmd/main.go`. Akkurat nå printer den bare Hello World. Vi bruker `air`. Dette er et vertkøy som bygger og kjører appen på nytt hver gang du har gjort en endring.

Åpne en ny tab i terminalen og kjør `air`. Den burde printe litt logs og så "Hello world" fulgt av Exit 0.

### Før du går videre

1. Air burde kjøre i bakgrunnen (IKKE stopp den).
2. Du får Hello world printet ut.

## 1. En enkel http handler

Siden vi lager et web API vil adapteren vår være en HTTP server. Altså, "inngangen" til logikken vår er gjennom HTTP requests. Alle adaptere vil befinne seg i en `/adapter` mappe. Vi skal lage en HTTP adapter, så den vil ligge under `/adapter/http`.

Lag en fil `ping.go`. Denne skal inneholde en handler som bare svarer med `Pong!` til endepunktet `/ping`.

Handlere har denne signaturen:

```go
func fooHandler(ctx *fuse.Context) error {
	// ...
}
```

Legg til denne enkle serveren i `/adapter/http/server.go` og mount ping handleren:

```go
func RunServer() {
	mux := fuse.NewMux(
		fuse.LoggerMiddleware(fuse.BasicLogger{}),
	)

	// Mount handlers her
	// ...
	
	http.ListenAndServe(":8080", mux)
}
```

Husk å kalle `RunServer()` i `main()`!

### Før du går videre

Test om endepunktet ditt fungerer som det skal. Du burde få "Pong!" som respons:

```sh
curl localhost:8080/ping
```

## 2. Liste alle ledige hoteller

Neste steg er å faktisk implementere et endepunkt. Start med følgende:

```sh
# Svarer med et JSON array av hoteller som er åpne og har ledige rom.
GET /hotels
```

### Steg

1. Legg til følgende type i `/domain/model/hotel.go`. Denne brukes av frontenden:

```go
type Hotel struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Stars        int    `json:"stars"`
	HasBreakfast bool   `json:"hasBreakfast"`
	HasGym       bool   `json:"hasGym"`
	HasPool      bool   `json:"hasPool"`
	Rooms        []Room `json:"rooms"`
}

type Room struct {
	RoomNumber string `json:"roomNumber"`
	Price      int    `json:"price"`
	Capacity   int    `json:"capacity"`
	Occupied   bool   `json:"occupied"`
}
```

> `json:"foo"` er en tag som spesifiserer hvilke JSON-felt navn som skal brukes for hvert felt i structen. Her gjør vi bare alt om til camelCase.

2. I domenelaget (`/domain/port/hotel.go`), lag et `interface` `HotelRepo` med en metode for å hente ut en liste av alle hoteller.
3. Implementer dette i infrastruktur-laget `/infra/hotelapi/hotel.go`.
4. I domenelaget igjen, lag en `struct` `HotelService` som har en `HotelRepo`. Lag en metode for å hente ut alle *ledige og åpne* hoteller.
5. Lag en handler i `adapter/http/hotel.go` som bruker en `HotelService` til å hente ut en liste med hoteller og svare i JSON.
6. Hook opp resten: initialiser `HotelRepo` og `HotelService` i main og gi `HotelService` til server funksjonen hvor handleren blir mountet.

Legg til denne statiske fil-handleren for å kunne se frontenden på `localhost:8080`:

```go
mux.Handle("GET", "/", func(ctx *fuse.Context) error {
	return ctx.ServeFile("web/index.html")
})
```

### Før du går videre

1. Du har et endepunkt `GET /hotels` som svarer med en JSON-liste med hoteller.
2. Alle hoteller i listen er åpne og har ledige rom.
3. Du kan se frontended på `localhost:8080` og får opp en liste med hoteller og rom.

## 3. Booke et hotellrom

Lag en ny handler for endepunktet `POST /hotels/{id}/room/{roomNr}` som lar deg booke et rom.

Det er lov å bruke chat for å lage frontenden for denne oppgaven.

### Før du går videre

1. Du har lagt til et nytt endepunkt med alle nødvendige ledd i `port`, `service`, `infra` osv.
2. Når du booker et hotell i frontenden vil statusen endre seg til Occupied.

## 4. Bli kjent med Uno

Denne oppgaven er en veldig forenklet versjon av hvordan DDD-apper er strukturert, men kjerneprinsippene er de samme.

Neste oppgave er å bli bedre kjent med Uno. Der er ting strukturert på nesten helt lik måte som denne oppgaven, med noen få unntak:

1. Det er ingen `adapter` mappe, bare `http/...`. Uno har ingen andre adaptere så vi tenkte det er best å slippe unødvendige mappe-lag.
2. I stedet for å initialisere avhengigheter i main gjør Uno det i en `bootstrap/bootstrap.go` fil. Her lages bare alle servicer og repos og gis videre til serveren.
3. Uno bruker en `Mux` (multiplexer) per route i stedet for enkelte handlers. Dette er bare for å kombinere ting og holde ting ryddig.
4. Pakken `fuse` som ble brukt i denne oppgaven er en forenklet versjon av det som brukes i Uno, men det aller meste er helt likt.
5. Uno bruker Swagger. Dette er en generert frontend som lar deg teste endepunkter. De rare kommentarene over handlers i `/http/routes/...` er det som generer specen som brukes av Swagger.

Det er sikkert mye annet jeg ikke kommer på...

### Steg

1. Lag en ny `HappeningWithRegistrationCount` type i domenet. Denne skal se slik ut:

```go
type HappeningWithRegistrationCount struct {
  Slug              string
  RegistrationCount int
}
```

2. Legg til nødvendige metoder i `HappeningRepo` og `HappeningService` for å hente ut denne dataen.
3. Lag en handler for endepunktet `GET /happenings/{id}/withRegCount` som returnerer denne typen som JSON for en gitt happening.

Det er mange detaljer som ikke blir nevnt i denne deloppgaven, men det er litt av poenget. Du burde kunne lese resten av Uno-koden og finne ut av hvordan ting gjøres. Det er VELDIG mye repetesjon, så det er helt lov (og anbefalt) å copy-paste annen kode og så endre på det litt.

### Du er ferdig når

1. Du har implementert endepunktet som beskrevet.
2. Du kan teste det i Swagger UI på `/api/swagger/index.html`.
