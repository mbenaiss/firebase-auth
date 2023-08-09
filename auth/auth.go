package auth

import "context"

// AuthenticationResult is the response from auth method.
type AuthenticationResult struct {
	ID           string `json:"localId"`
	Token        string `json:"idToken"`
	Email        string `json:"email"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
}

// Auth is an interface for authentication methods.
type Auth interface {
	SignUp(ctx context.Context, email, password string) (AuthenticationResult, error)
	SignIn(ctx context.Context, email, password string) (AuthenticationResult, error)
	ValidateToken(ctx context.Context, token string) (string, error)
	DeleteUser(ctx context.Context, uid string) error
	// ExistByEmail check if the user exists and returns the uid of the user.
	ExistByEmail(ctx context.Context, email string) (string, error)
	// UpdatePassword update the password of the user.
	UpdatePassword(ctx context.Context, uid, password string) error
	// UpdateEmail update the email of the user.
	UpdateEmail(ctx context.Context, uid, email string) error
	// SetCustomUserClaims set custom claims for the user.
	SetCustomUserClaims(ctx context.Context, uid string, claims map[string]interface{}) error
	// GetCustomUserClaims get custom claims for the user.
	GetUserClaims(ctx context.Context, uid string) (map[string]interface{}, error)
}
