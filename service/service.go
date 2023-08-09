package service

import "github.com/auth/auth"

// Service is a struct representing a Service.
type Service struct {
	auth auth.Auth
}

// New creates a new Service instance.
func New(auth auth.Auth) *Service {
	return &Service{
		auth: auth,
	}
}
