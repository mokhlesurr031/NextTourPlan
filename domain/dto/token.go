package dto

import (
	"time"
)

type LoggerInUserData struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
	Contact  string `json:"contact"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	Rating   int    `json:"rating"`
}
type JWTToken struct {
	ExpiredIn time.Duration
	MaxAge    int
	Secret    string
	Message   string
	User      *LoggerInUserData
}
