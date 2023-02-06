package service

import (
	"WEB_REST_exm0302"
	"WEB_REST_exm0302/pkg/repository"
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	salt       = "asdjfbvlkdfnbjlvcxkjkbu" //для хэширования
	signingKey = "srfjtfryjfgnbmnH"
	tokenTTL   = 12 * time.Hour
)

// дополнить стандартный claims с id пользователя
type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

// В конструкторе принимаем репозиторий с базой
func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

// Передаем структуру в репозиторий
func (s *AuthService) CreateUser(user WEB_REST_exm0302.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

// Получение пользователя из БД
func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

// хэширование пароля
func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
