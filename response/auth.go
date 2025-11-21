package response

import (
	"time"
)

type Login struct {
	Token string `json:"token"`
	TokenType string `json:"token_type"`
	ExpiresIn int `json:"expires_in"`
}

type TokenIntrospect struct {
	ExpireAt time.Time `json:"expire_at"`
	ID string `json:"id"`
	Name string `json:"name"`
	Email	string `json:"email"`
}