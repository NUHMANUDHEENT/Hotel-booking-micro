package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"

	"github.com/dgrijalva/jwt-go"
	"github.com/nuhmanudheent/hotel-booking/user-service/internal/domain"
	"github.com/nuhmanudheent/hotel-booking/user-service/internal/repository"

	middleware "github.com/nuhmanudheent/hotel-booking/user-service/pkg/middlerware-http"
	hotel_service "github.com/nuhmanudheent/hotel-booking/user-service/proto/client_proto"
)

type UserService interface {
	RegisterUser(name, email, password, mobile string) error
	LoginUser(email, password string) (string, error)
	UserGetInfo(id uint) (domain.User, error)
	GetHotelRooms() ([]*hotel_service.Room, error)
	UserExists(userID uint32) bool
}

type userService struct {
	repo        repository.UserRepository
	hotelClient hotel_service.HotelServiceClient
}

func NewUserService(repo repository.UserRepository, hotelConn *grpc.ClientConn) UserService {
	hotelClient := hotel_service.NewHotelServiceClient(hotelConn)
	return &userService{repo: repo, hotelClient: hotelClient}
}
func (u *userService) RegisterUser(name, email, password, mobile string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := domain.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		Mobile:   mobile,
	}

	err = u.repo.RegisterUser(user)
	return err
}

func (u *userService) LoginUser(email, password string) (string, error) {
	user, err := u.repo.FindByEmail(email)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid email or password")
	}
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &middleware.Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(middleware.JwtKey)
	if err != nil {
		return "", err
	}
	fmt.Println("jwt", user.ID)
	return tokenString, nil
}
func (u *userService) UserGetInfo(id uint) (domain.User, error) {
	user, err := u.repo.FindByID(id)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
func (u *userService) GetHotelRooms() ([]*hotel_service.Room, error) {
	req := &hotel_service.GetRoomsRequest{}
	resp, err := u.hotelClient.GetRooms(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return resp.Rooms, nil
}
func (s *userService) UserExists(userID uint32) bool {
	return s.repo.CheckUser(userID)
}
