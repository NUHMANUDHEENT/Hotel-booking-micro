package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nuhmanudheent/hotel-booking/booking-service/internal/domain"
	"github.com/nuhmanudheent/hotel-booking/booking-service/internal/service"
	http_middleware "github.com/nuhmanudheent/hotel-booking/booking-service/pkg/middlerware_http"
)

type BookingHandler struct {
	bookingService service.BookingService
}

func NewBookingHandler(bookingService service.BookingService) *BookingHandler {
	return &BookingHandler{bookingService: bookingService}
}

func (h *BookingHandler) BookingRouters(r *gin.Engine) {
	r.POST("/book", http_middleware.AuthMiddleware(), h.CreateBooking)
}

func (h *BookingHandler) CreateBooking(c *gin.Context) {
	var req domain.Booking
	userId := c.GetUint("userID")
	// Bind JSON request to the struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	// // Parse timestamps
	// checkInTime, err := time.Parse(time.RFC3339, req.CheckIn)
	// if err != nil {
	//     c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid check-in time format"})
	//     return
	// }

	// checkOutTime, err := time.Parse(time.RFC3339, req.CheckOut)
	// if err != nil {
	//     c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid check-out time format"})
	//     return
	// }

	// Call service function

	booking, err := h.bookingService.CreateBooking(uint32(userId), req.RoomID, req.CheckIn, req.CheckOut, req.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return response
	c.JSON(http.StatusOK, gin.H{
		"booking_id":     booking.OrderId,
		"razor_order_id": booking.RazorOrderId,
		"status":         booking.Status,
	})
}
