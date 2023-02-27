package user

import "golang.org/x/crypto/bcrypt"

func Encrypt(passwd *string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(*passwd), 10)
	return string(hashed)
}
