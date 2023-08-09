package models

import (
	"fmt"
	"time"
)

// SignInInput is the request input for SignIn.
type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SignUpInput is the request input for SignUp.
type SignUpInput struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"user_name"`
}

// Validate validates the signup fields.
func (s *SignUpInput) Validate() error {
	if s.Email == "" {
		return fmt.Errorf("email is required")
	}

	return nil
}

// PasswordReset is a struct that holds the data for a password reset.
type PasswordReset struct {
	ID        string    `json:"id" firestore:"id"`
	CreatedAt time.Time `json:"created_at" firestore:"created_at"`
	Email     string    `json:"email" firestore:"email"`
	Token     string    `json:"token" firestore:"token"`
}

// PasswordResetInput is the request input for PasswordReset.
type PasswordResetInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token" firestore:"token"`
}

// AuthResult is the response from auth method.
type AuthResult struct {
	ID           string `json:"localId"`
	Token        string `json:"idToken"`
	Email        string `json:"email"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
}
