package repository

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nuhmanudheent/hotel-booking/payment-service/internal/domain"
	"github.com/razorpay/razorpay-go"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	PaymentComplete(c *gin.Context)
	PaymentConfirmation(c *gin.Context) bool
	PaymentCheck(orderId string) (string, error)
	NewRazorOrder(orderId string, price uint32) (string, error)
}
type paymentRepostory struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepostory{db: db}
}
func (p *paymentRepostory) PaymentComplete(c *gin.Context) {
	c.HTML(200, "payment.html", nil)
}
func (p *paymentRepostory) PaymentCheck(orderId string) (string, error) {
	var paymentcheck domain.PaymentDetails
	if err := p.db.First(&paymentcheck, "order_id= ?", orderId).Error; err != nil {
		return "", errors.New("failed to get payment details")
	}
	return paymentcheck.Status, nil
}

func (p *paymentRepostory) PaymentConfirmation(c *gin.Context)(string, string) {
	var paymentStore domain.PaymentDetails
	var paymentDetails = make(map[string]string)
	if err := c.BindJSON(&paymentDetails); err != nil {
		log.Fatal("failed to fetch payment data")
	}
	pd := paymentDetails
	//============== verify the signature ================
	err := RazorPaymentVerification(pd["signature"], pd["order_id"], pd["payment_id"])
	if err != nil {
		log.Fatal("-------", err)
	}
	if err := p.db.First(&paymentStore, "order_id=?", pd["order_id"]).Error; err != nil {
		log.Fatal("can't find payment details")
	}
	paymentStore.PaymentId = pd["payment_id"]
	paymentStore.Status = "success"
	p.db.Create(&paymentStore)

	return paymentStore.,paymentStore.Status
}

func RazorPaymentVerification(sign, orderId, paymentId string) error {
	signature := sign
	secret := os.Getenv("RAZOR_PAY_SECRET")
	data := orderId + "|" + paymentId
	h := hmac.New(sha256.New, []byte(secret))
	_, err := h.Write([]byte(data))
	if err != nil {
		panic(err)
	}
	sha := hex.EncodeToString(h.Sum(nil))
	if subtle.ConstantTimeCompare([]byte(sha), []byte(signature)) != 1 {
		return errors.New("PAYMENT FAILED")
	} else {
		return nil
	}
}
func (p *paymentRepostory) NewRazorOrder(orderId string, price uint32) (string, error) {
	client := razorpay.NewClient(os.Getenv("RAZOR_PAY_KEY"), os.Getenv("RAZOR_PAY_SECRET"))
	orderParams := map[string]interface{}{
		"amount":   price * 100,
		"currency": "INR",
		"receipt":  orderId,
	}
	order, err := client.Order.Create(orderParams, nil)
	if err != nil {
		return "", errors.New("PAYMENT NOT INITIATED")
	}

	razorId, _ := order["id"].(string)
	return razorId, nil
}
