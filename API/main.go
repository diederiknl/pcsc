package API

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"time"
)

// Struct voor inkomende data.
// ID heb ik afgevinkt. Volgens mij hebben we geen noodzaak voor een ID.

type SmartCardData struct {
	//	ID        uint      `gorm:"primaryKey"`
	ReaderID  string    `json:"reader_id"`
	CardID    string    `json:"card_id"`
	Timestamp time.Time `json:"timestamp"`
}

// Pointer naar de database.
var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open(sqlite.Open("smartcard.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&SmartCardData{})

	r := gin.Default()
	r.POST("/data", postData)
	r.Run(":8080")
}

func postData(c *gin.Context) {
	var input SmartCardData
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.Timestamp = time.Now()
	db.Create(&input)

	c.JSON(http.StatusOK, gin.H{"data": input})
}
