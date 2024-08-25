package middleware

import "golang.org/x/crypto/bcrypt"

func Bcrypt(password string) string {
	salt := 3
	pwd := []byte(password)
	hash, _ := bcrypt.GenerateFromPassword(pwd, salt)
	return string(hash)
}

func CheckPasswordHash(h, p []byte) bool {
	hash, pass := []byte(h), []byte(p)
	err := bcrypt.CompareHashAndPassword(hash, pass)
	return err == nil
}
