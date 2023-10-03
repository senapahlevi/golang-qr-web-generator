package generateqr

import (
	"net/http"
	"qrweb/databases"
	"qrweb/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetQRCodeController(databases *databases.DB) {
	db = databases.DB
}

func SetDatabaseRegister(databases *databases.DB) {
	db = databases.DB
}
func GenerateQRCode(c *gin.Context) {

	var data models.QRcode
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// qrcodeImage, err := qrcode.Encode(data.Text, qrcode.Medium, 256)
	qrcodeImage, err := qrcode.Encode("tel:"+data.Text, qrcode.Medium, 256)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	qrcodes := models.QRcode{
		Text:      data.Text,
		TextURL:   data.TextURL,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	db.Create(&qrcodes)

	c.Data(http.StatusOK, "image/png", qrcodeImage)

}

type QRcodeRequest struct {
	Text string `gorm:"text" json:"text"`
}
