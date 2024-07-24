package http

import (
	"github.com/gin-gonic/gin"
	"github.com/nuhmanudheent/hotel-booking/payment-service/internal/service"
)

type PaymentHandler struct {
	paymentService service.PaymentService
}

func NewPaymentHandler(paymentService service.PaymentService) *PaymentHandler {
	return &PaymentHandler{paymentService: paymentService}
}

func (h *PaymentHandler) RegisterRoutes(r *gin.Engine) {
	r.GET("/payment/complete", h.PaymentComplete)
	r.POST("/payment/confirm", h.PaymentConfirmation)
}

func (h *PaymentHandler) PaymentComplete(c *gin.Context) {
	h.paymentService.PaymentComplete(c)
}

func (h *PaymentHandler) PaymentConfirmation(c *gin.Context) {
	h.paymentService.PaymentConfirmation(c)
}
