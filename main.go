package main

import (
	"fmt"
	"log"

	"github.com/deeper-x/gopcsc/smartcard"
)

func main() {
	// Establish context
	ctx, err := smartcard.EstablishContext()
	if err != nil {
		log.Fatalf("Failed to establish context: %v", err)
	}
	defer ctx.Release()

	// Wait for card to be present
	reader, err := ctx.WaitForCardPresent()
	if err != nil {
		log.Fatalf("Failed to wait for card present: %v", err)
	}

	// Connect to the card
	card, err := reader.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to card: %v", err)
	}
	defer card.Disconnect()

	// Print the ATR of the card
	fmt.Printf("Card ATR: %X\n", card.ATR())

	// Create APDU command (SelectCommand function)
	command := SelectCommand(0xa0, 0x00, 0x00, 0x00, 0x62, 0x03, 0x01, 0x0c, 0x01, 0x01)

	// Transmit APDU command
	response, err := card.TransmitAPDU(command)
	if err != nil {
		log.Fatalf("Failed to transmit APDU: %v", err)
	}

	// Print response
	fmt.Printf("Response: %X\n", response)
}

// SelectCommand constructs an APDU command
func SelectCommand(cla, ins, p1, p2 byte, data ...byte) []byte {
	lc := byte(len(data))
	apdu := append([]byte{cla, ins, p1, p2, lc}, data...)
	return apdu
}
