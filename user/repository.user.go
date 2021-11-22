package book

import (
	"log"
	"pustaka-api/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type IUserRepository interface {
	InsertUser(user models.User) (models.User)
	// UpdateUser(user models.User) models.User
	// VerifyCredential(email string, password string) interface{}
	// IsDuplicateEmail(email string) (tx *gorm.DB)
	// FindByEmail(email string) models.User
	// ProfileUser(userID string) models.User
}

type PustakaApiRepository struct {
	pustaka_api *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &PustakaApiRepository{
		pustaka_api: db,
	}
}

func (r *PustakaApiRepository) InsertUser(user models.User) models.User {
	user.Password = hashAndSalt([]byte(user.Password))
	r.pustaka_api.Save(&user)

	return user
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}
