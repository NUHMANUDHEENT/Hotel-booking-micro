package http

import (
  "github.com/gin-gonic/gin"
  "github.com/nuhmanudheent/hotel_booking/hotel_service/internal/domain"
  "github.com/nuhmanudheent/hotel_booking/hotel_service/internal/service"
  "net/http"
)

type HotelHandler struct {
  hotelService service.HotelService
}

func NewHotelHandler(hotelService service.HotelService) *HotelHandler {
  return &HotelHandler{hotelService: hotelService}
}

func (h *HotelHandler) HotelRouters(r *gin.Engine) {
  r.POST("add", h.AddRoom)
  r.GET("get/:roomNumber", h.GetRoom)
  r.PUT("update", h.UpdateRoom)
  r.DELETE("delete/:roomNumber", h.DeleteRoom)
}

func (h *HotelHandler) AddRoom(c *gin.Context) {
  var room domain.Room
  if err := c.ShouldBindJSON(&room); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
    return
  }
  err := h.hotelService.AddRoom(room)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, gin.H{"message": "Room added successfully"})
}

func (h *HotelHandler) GetRoom(c *gin.Context) {
  roomNumber := c.Param("roomNumber")
  room, err := h.hotelService.GetRoom(roomNumber)
  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, gin.H{"room": room})
}

func (h *HotelHandler) UpdateRoom(c *gin.Context) {
  var room domain.Room
  if err := c.ShouldBindJSON(&room); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
    return
  }
  err := h.hotelService.UpdateRoom(room)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, gin.H{"message": "Room updated successfully"})
}

func (h *HotelHandler) DeleteRoom(c *gin.Context) {
  roomNumber := c.Param("roomNumber")
  err := h.hotelService.DeleteRoom(roomNumber)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, gin.H{"message": "Room deleted successfully"})
}
