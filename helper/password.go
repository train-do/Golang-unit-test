package helper

import "golang.org/x/crypto/bcrypt"

func CheckPassword(inputPassword, storedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(inputPassword))
	return err == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
