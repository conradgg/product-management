package main

import "os"

func LoadConfig() {
	os.Setenv("DATABASE_URL", "postgres://postgres:postgrespw@localhost:32770/test")
}
