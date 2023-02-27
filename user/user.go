package user

import (
	"fmt"
	"product-management/postgresql"
)

type User struct {
	UserName  string
	FirstName string
	LastName  string
	Password  string
}

func UserExist(user *string) (exist bool) {
	q := fmt.Sprintf("SELECT username FROM users WHERE username='%s'", *user)
	if postgresql.QueryRow(&q) != "" {
		exist = true
	}
	return
}

func Register(u *User) {
	q := fmt.Sprintf("INSERT INTO users (username, firstname, lastname, password) VALUES ('%s', '%s', '%s', '%s')", u.UserName, u.FirstName, u.LastName, u.Password)
	postgresql.Exec(&q)
}

func Auth(u *User) (pass bool) {
	q := fmt.Sprintf("SELECT * FROM users WHERE username='%s' AND password='%s'", u.UserName, u.Password)
	if len(postgresql.Query(&q)) == 1 {
		pass = true
	}
	return
}
