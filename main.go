package main

func main() {
	// Lees de smartcard uit
	data := pkg.ReadSmartCard()

	// Stuur de data naar de REST-API
	pkg.SendToAPI(data)
}
