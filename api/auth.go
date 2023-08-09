package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/auth/models"
	"github.com/auth/service"
)

func signin(svc *service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body models.SignInInput

		err := c.BindJSON(&body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		resp, err := svc.SignIn(c.Request.Context(), body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

			return
		}

		c.JSON(http.StatusOK, resp)
	}
}

func signup(svc *service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var signupInput models.SignUpInput

		err := c.BindJSON(&signupInput)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		memberID := c.Query("member_id")
		lodgerID := c.Query("lodger_id")

		err = signupInput.Validate()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		resp, err := svc.SignUp(c.Request.Context(), signupInput, memberID, lodgerID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

			return
		}

		c.JSON(http.StatusOK, resp)
	}
}

func resetPassword(svc *service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body models.PasswordResetInput

		err := c.BindJSON(&body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		err = svc.ResetPassword(c.Request.Context(), body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

			return
		}

		c.Status(http.StatusOK)
	}
}

func validateToken(svc *service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body models.PasswordResetInput

		err := c.BindJSON(&body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		err = svc.ValidateToken(c.Request.Context(), body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

			return
		}

		c.Status(http.StatusOK)
	}
}
