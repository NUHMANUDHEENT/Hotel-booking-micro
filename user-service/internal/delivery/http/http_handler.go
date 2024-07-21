package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nuhmanudheent/hotel-booking/user-service/internal/domain"
	"github.com/nuhmanudheent/hotel-booking/user-service/internal/service"
	middleware "github.com/nuhmanudheent/hotel-booking/user-service/pkg/middlerware-http"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (u *UserHandler) UserRouters(r *gin.Engine) {
	r.POST("register", u.RegisterUser)
	r.POST("login", u.LoginUser)
	r.GET("profile", middleware.AuthMiddleware(), u.UserGetInfo)
	r.GET("hotels", middleware.AuthMiddleware(), u.GetHotelRooms)
}

func (u *UserHandler) RegisterUser(c *gin.Context) {
	var userReg domain.User
	if err := c.ShouldBindJSON(&userReg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	err := u.userService.RegisterUser(userReg.Name, userReg.Email, userReg.Password, userReg.Mobile)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
	})
}

func (u *UserHandler) LoginUser(c *gin.Context) {
	var loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	user, err := u.userService.LoginUser(loginRequest.Email, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"jwt":     user,
	})
}

func (u *UserHandler) UserGetInfo(c *gin.Context) {
	userId := c.GetUint("userID")
	fmt.Println("userrr", userId)
	user, err := u.userService.UserGetInfo(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	userData := gin.H{
		"name":   user.Name,
		"email":  user.Email,
		"mobile": user.Mobile,
	}
	c.JSON(http.StatusOK, userData)
}

func (u *UserHandler) GetHotelRooms(c *gin.Context) {
	rooms, err := u.userService.GetHotelRooms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get hotel rooms", "message": err.Error()})
		return
	}

	var roomData []gin.H
	for _, room := range rooms {
		roomData = append(roomData, gin.H{
			"room_number": room.RoomNumber,
			"category":    room.Category,
			"name":        room.Availability,
			"price":       room.Price,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"rooms": roomData,
	})
}
