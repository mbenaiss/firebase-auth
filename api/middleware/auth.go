package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type uidContextKey string

const (
	bearer = "Bearer "

	// UIDKey is the key for the uid of the current user in the context.
	UIDKey uidContextKey = "uid"
)

type auth interface {
	ValidateToken(ctx context.Context, token string) (string, error)
}

// Authorization verifies the token provided is Authorization header.
func Authorization(auth auth) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		if authorization == "" || !strings.HasPrefix(authorization, bearer) {
			c.AbortWithStatus(http.StatusUnauthorized)

			return
		}

		token := strings.TrimPrefix(authorization, bearer)

		uid, err := auth.ValidateToken(c.Request.Context(), token)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)

			return
		}

		ctx := context.WithValue(c.Request.Context(), UIDKey, uid)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
