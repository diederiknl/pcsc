package awe_cardreader

import (
	"encoding/json"
	"fmt"
	"github.com/ebfe/scard"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type SmartCardData struct {
	Data      string    `json:"data"`
	Timestamp time.Time `json:"timestamp"`
}

func ReadSmartCard() string {
	ctx, err := scard.EstablishContext()
	if err != nil {
		fmt.Println("Error establishing return :", err)
	}
	defer ctx.Release()

	readers, err := ctx.ListReaders()
	if err != nil {
		fmt.Println("Error listing readers:", err)

	}

	if len(readers) == 0 {
		fmt.Println("No smartcard readers found")

	}

	reader := readers[0]
	card, err := ctx.Connect(reader, scard.ShareShared, scard.ProtocolAny)
	if err != nil {
		fmt.Println("Error connecting to card:", err)

	}
	defer card.Disconnect(scard.LeaveCard)

	cmd := []byte{0x00, 0xA4, 0x04, 0x00, 0x00} // Example APDU command
	resp, err := card.Transmit(cmd)
	if err != nil {
		fmt.Println("Error transmitting APDU:", err)

	}

	fmt.Printf("Response: %X\n", resp)
	return string(resp)
}

func SendToAPI(data string) {
	url := "http://localhost:8080/data"
	smartCardData := SmartCardData{
		Data:      data,
		Timestamp: time.Now(),
	}

	jsonData, err := json.Marshal(smartCardData)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	req, err := http.NewRequest("POST", url, ioutil.NopCloser(strings.NewReader(string(jsonData))))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	fmt.Println("Response from API:", string(body))
}
