package user

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	result := db.Create(user)
	if err := result.Error; err != nil {
		return err
	}

	return nil
}

func GetUser(db *gorm.DB, email string) (*User, error) {
	u := new(User)
	result := db.Where("email = ?", email).First(u)
	if err := result.Error; err != nil {
		return nil, err
	}

	return u, nil
}
