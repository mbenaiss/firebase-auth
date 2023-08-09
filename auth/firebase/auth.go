package firebaseauth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	firebase "firebase.google.com/go"
	firebaseauth "firebase.google.com/go/auth"
	"github.com/auth/auth"
)

const googleAPIURL = "https://identitytoolkit.googleapis.com/v1/accounts"

type firebaseAuth struct {
	httpClient *http.Client
	apiKey     string
	auth       *firebaseauth.Client
}

// New creates a new Auth instance.
func New(ctx context.Context, apiKey string, app *firebase.App) (auth.Auth, error) {
	auth, err := app.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to init firebase auth %w", err)
	}

	return &firebaseAuth{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		apiKey: apiKey,
		auth:   auth,
	}, nil
}

type errorResponse struct {
	Error err `json:"error"`
}

type err struct {
	Message string `json:"message"`
}

type apiRequest struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ReturnSecureToken bool   `json:"returnSecureToken"`
}

// SignUp creates a new user and returns the id and the token for the user.
func (a *firebaseAuth) SignUp(ctx context.Context, email, password string) (auth.AuthenticationResult, error) {
	input := apiRequest{
		Email:             email,
		Password:          password,
		ReturnSecureToken: true,
	}

	buf, err := json.Marshal(&input)
	if err != nil {
		return auth.AuthenticationResult{}, fmt.Errorf("unable to marshal input to signup new user: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, googleAPIURL+":signUp", bytes.NewBuffer(buf))
	if err != nil {
		return auth.AuthenticationResult{}, fmt.Errorf("unable to init request to signup new user: %w", err)
	}

	values := req.URL.Query()
	values.Add("key", a.apiKey)
	req.URL.RawQuery = values.Encode()

	req.Header.Add("Content-Type", "application/json")

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return auth.AuthenticationResult{}, fmt.Errorf("unable to make request to signup new user: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errBody errorResponse

		err = json.NewDecoder(resp.Body).Decode(&errBody)
		if err != nil {
			return auth.AuthenticationResult{}, fmt.Errorf("unable to decode error response to signup: %w", err)
		}

		return auth.AuthenticationResult{}, fmt.Errorf("unable to signup new user: %s", errBody.Error.Message)
	}

	var r auth.AuthenticationResult

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return auth.AuthenticationResult{}, fmt.Errorf("unable to decode response to signup new user: %w", err)
	}

	return r, nil
}

// SignIn validate the provided credentials and returns the token.
func (a *firebaseAuth) SignIn(ctx context.Context, email, password string) (auth.AuthenticationResult, error) {
	input := apiRequest{
		Email:             email,
		Password:          password,
		ReturnSecureToken: true,
	}

	buf, err := json.Marshal(&input)
	if err != nil {
		return auth.AuthenticationResult{}, fmt.Errorf("unable to marshal input to signin user: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, googleAPIURL+":signInWithPassword", bytes.NewBuffer(buf))
	if err != nil {
		return auth.AuthenticationResult{}, fmt.Errorf("unable to init request to signin user: %w", err)
	}

	values := req.URL.Query()
	values.Add("key", a.apiKey)
	req.URL.RawQuery = values.Encode()

	req.Header.Add("Content-Type", "application/json")

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return auth.AuthenticationResult{}, fmt.Errorf("unable to make request to signin user: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errBody errorResponse

		err = json.NewDecoder(resp.Body).Decode(&errBody)
		if err != nil {
			return auth.AuthenticationResult{}, fmt.Errorf("unable to decode error response to signin: %w", err)
		}

		return auth.AuthenticationResult{}, fmt.Errorf("unable to signin the user: %s", errBody.Error.Message)
	}

	var r auth.AuthenticationResult

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return auth.AuthenticationResult{}, fmt.Errorf("unable to decode response to signin user: %w", err)
	}

	return r, nil
}

// ValidateToken validate the provided token and returns the uid of the user.
func (a *firebaseAuth) ValidateToken(ctx context.Context, token string) (string, error) {
	decodedToken, err := a.auth.VerifyIDToken(ctx, token)
	if err != nil {
		return "", fmt.Errorf("unable to validate token: %w", err)
	}

	return decodedToken.UID, nil
}

// DeleteUser delete user from firebase auth.
func (a *firebaseAuth) DeleteUser(ctx context.Context, uid string) error {
	err := a.auth.DeleteUser(ctx, uid)
	if err != nil {
		return fmt.Errorf("unable to delete user with uid %s: %w", uid, err)
	}

	return nil
}

// ExistByEmail check if the user exists and returns the uid of the user.
func (a *firebaseAuth) ExistByEmail(ctx context.Context, email string) (string, error) {
	user, err := a.auth.GetUserByEmail(ctx, email)
	if err != nil {
		return "", fmt.Errorf("unable to find user by email %s: %w", email, err)
	}

	return user.UID, nil
}

// UpdatePassword update the password of the user.
func (a *firebaseAuth) UpdatePassword(ctx context.Context, uid, password string) error {
	user := &firebaseauth.UserToUpdate{}
	user.Password(password)

	_, err := a.auth.UpdateUser(ctx, uid, user)
	if err != nil {
		return fmt.Errorf("unable to update user: %w", err)
	}

	return nil
}

// UpdateEmail update the email of the user.
func (a *firebaseAuth) UpdateEmail(ctx context.Context, uid, email string) error {
	user := &firebaseauth.UserToUpdate{}
	user.Email(email)

	_, err := a.auth.UpdateUser(ctx, uid, user)
	if err != nil {
		return fmt.Errorf("unable to update user email: %w", err)
	}

	return nil
}

// SetCustomUserClaims set custom claims for the user.
func (a *firebaseAuth) SetCustomUserClaims(ctx context.Context, uid string, claims map[string]interface{}) error {
	err := a.auth.SetCustomUserClaims(ctx, uid, claims)
	if err != nil {
		return fmt.Errorf("unable to set custom claims for user_id %s: %w", uid, err)
	}

	return nil
}

// GetCustomUserClaims get custom claims for the user.
func (a *firebaseAuth) GetUserClaims(ctx context.Context, uid string) (map[string]interface{}, error) {
	claims, err := a.auth.GetUser(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("unable to get user claims for user_id %s: %w", uid, err)
	}

	return claims.CustomClaims, nil
}
