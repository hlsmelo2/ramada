package auth

import "golang.org/x/crypto/bcrypt"

func HashIt(toCrypt string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(toCrypt), bcrypt.DefaultCost)
}

func CheckPassword(password, hashed string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
}
