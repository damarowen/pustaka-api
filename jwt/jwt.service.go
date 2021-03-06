package JWT

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

//JWTService is a contract of what jwtService can do
type IJwtService interface {
	GenerateToken(userID string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	issuer    string
	secretKey string
}

//NewJWTService method is creates a new instance of JWTService
func NewJWTService() IJwtService {
	return &jwtService{
		issuer:    "belajargolang",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey != "" {
		secretKey = "belajargolang"
	}
	return secretKey
}

func (j *jwtService) GenerateToken(UserID string) string {
	claims := &jwtCustomClaim{
		UserID,
		jwt.StandardClaims{
			//* addDate year, month, day = 1 tahun
			// ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			//* expired 1 minute
			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	//* algoritma 526
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}