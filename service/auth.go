package service

import (
	"context"
	"fmt"

	"github.com/auth/models"
)

// SignIn validate the provided credentials and returns the token.
func (s *Service) SignIn(ctx context.Context, input models.SignInInput) (models.AuthResult, error) {
	res, err := s.auth.SignIn(ctx, input.Email, input.Password)
	if err != nil {
		return models.AuthResult{}, fmt.Errorf("failed to sign in %w", err)
	}

	authResp := models.AuthResult{
		ID:           res.ID,
		Token:        res.Token,
		Email:        res.Email,
		RefreshToken: res.RefreshToken,
		ExpiresIn:    res.RefreshToken,
	}

	return authResp, nil
}

// SignUp creates a new user and returns the id and the token for the user.
func (s *Service) SignUp(ctx context.Context, input models.SignUpInput, memberID string, lodgerID string) (models.AuthResult, error) {
	authenticationResult, err := s.auth.SignUp(ctx, input.Email, input.Password)
	if err != nil {
		return models.AuthResult{}, fmt.Errorf("failed to create user %w", err)
	}

	return models.AuthResult{
		ID:           authenticationResult.ID,
		Token:        authenticationResult.Token,
		Email:        authenticationResult.Email,
		RefreshToken: authenticationResult.RefreshToken,
		ExpiresIn:    authenticationResult.RefreshToken,
	}, nil
}

// ValidateToken validates the token.
func (s *Service) ValidateToken(ctx context.Context, input models.PasswordResetInput) error {
	_, err := s.auth.ValidateToken(ctx, input.Token)
	if err != nil {
		return fmt.Errorf("invalid token %w", err)
	}

	return nil
}

// ResetPassword resets the password of the user.
func (s *Service) ResetPassword(ctx context.Context, input models.PasswordResetInput) error {
	_, err := s.auth.ValidateToken(ctx, input.Token)
	if err != nil {
		return fmt.Errorf("invalid token %w", err)
	}

	err = s.auth.UpdatePassword(ctx, input.Email, input.Password)
	if err != nil {
		return fmt.Errorf("failed to update password %w", err)
	}

	return nil
}
