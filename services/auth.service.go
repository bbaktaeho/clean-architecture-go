package services

import (
	"clean-architecture-go/dto"
	"clean-architecture-go/entities"
	"clean-architecture-go/repositories"
	"log"

	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(email, password string) interface{}
	CreateUser(user dto.RegisterDTO) entities.User
	FindByEmail(email string) entities.User
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepository repositories.UserRepository
}

func NewAuthService(userRepository repositories.UserRepository) AuthService {
	return &authService{
		userRepository,
	}
}

func (service *authService) VerifyCredential(email, password string) interface{} {
	res := service.userRepository.VerifyCredential(email, password)
	if v, ok := res.(entities.User); ok {
		comparedPassword := comparedPassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func (service *authService) CreateUser(user dto.RegisterDTO) entities.User {
	userToCreate := entities.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.userRepository.InsertUser(userToCreate)
	return res
}

func (service *authService) FindByEmail(email string) entities.User {
	return service.userRepository.FindByEmail(email)
}

func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func comparedPassword(hashedPw string, plainPw []byte) bool {
	byteHash := []byte(hashedPw)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPw)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
