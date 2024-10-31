package main

import (
	"github.com/diederiknl/PCSC/awe_cardreader"
)

func main() {
	// Lees de smartcard uit
	data := awe_cardreader.ReadSmartCard()

	// Stuur de data naar de REST-API
	awe_cardreader.SendToAPI(data)
}
