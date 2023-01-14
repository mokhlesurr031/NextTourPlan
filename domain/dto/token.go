package dto

import "time"

type JWTToken struct {
	ExpiredIn time.Duration
	MaxAge    int
	Secret    string
	Message   string
}
