package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/uma-universal-money-address/uma-go-sdk/uma/errors"
	"github.com/uma-universal-money-address/uma-go-sdk/uma/generated"
)

// User represents a user in the system.
type User struct {
	ID             string
	Name           string
	HashedPassword string
}

func (u User) GetUmaAddress(c *UmaConfig, context *gin.Context) string {
	leadingDollar := "$"
	if strings.HasPrefix(u.Name, leadingDollar) {
		leadingDollar = ""
	}
	return fmt.Sprintf("%s%s@%s", leadingDollar, u.Name, c.GetVaspDomain(context))
}

type UserService interface {
	// GetUser returns a user by ID.
	GetUser(id string) (*User, error)

	// GetUserFromContext Gets a user from a gin context.
	GetUserFromContext(context *gin.Context) (*User, error)

	// GetUserByUmaAddress returns a user by UMA address.
	GetUserByUmaAddress(umaAddress string, config UmaConfig, context *gin.Context) (*User, error)
}

// UserServiceFromEnv is a UserService that uses environment variables to get user information for a single user.
type UserServiceFromEnv struct {
	user *User
}

func NewUserServiceFromEnv(config UmaConfig) *UserServiceFromEnv {
	return &UserServiceFromEnv{
		user: &User{
			ID:             config.UserID,
			Name:           config.Username,
			HashedPassword: config.HashedUserPassword,
		},
	}
}

func (u *UserServiceFromEnv) GetUser(id string) (*User, error) {
	if u.user == nil || u.user.ID != id {
		return nil, &errors.UmaError{
			Reason:    "User not found",
			ErrorCode: generated.UserNotFound,
		}
	}
	return u.user, nil
}

func (u *UserServiceFromEnv) GetUserFromContext(context *gin.Context) (*User, error) {
	if u.user == nil {
		return nil, &errors.UmaError{
			Reason:    "User not found",
			ErrorCode: generated.UserNotFound,
		}
	}

	// Get from session cookie:
	session := sessions.Default(context)
	if session.Get("user_id") == u.user.ID {
		return u.user, nil
	}

	// Get from authorization header:
	authHeader := context.GetHeader("Authorization")
	basicToken := "Basic "
	if strings.HasPrefix(authHeader, basicToken) {
		encodedCredentials := strings.TrimPrefix(authHeader, basicToken)
		decodedCredentials, err := base64.StdEncoding.DecodeString(encodedCredentials)
		if err != nil {
			return nil, &errors.UmaError{
				Reason:    fmt.Sprintf("error decoding credentials: %v", err),
				ErrorCode: generated.Forbidden,
			}
		}
		credentials := strings.Split(string(decodedCredentials), ":")
		if len(credentials) != 2 {
			return nil, &errors.UmaError{
				Reason:    "invalid credentials",
				ErrorCode: generated.Forbidden,
			}
		}
		password := credentials[1]
		hashedPassword := hashString(password)
		if credentials[0] == u.user.Name && hashedPassword == u.user.HashedPassword {
			return u.user, nil
		}
	}

	return nil, &errors.UmaError{
		Reason:    "user not found",
		ErrorCode: generated.UserNotFound,
	}
}

func (u *UserServiceFromEnv) GetUserByUmaAddress(umaAddress string, config UmaConfig, context *gin.Context) (*User, error) {
	if u.user == nil || u.user.GetUmaAddress(&config, context) != umaAddress {
		return nil, &errors.UmaError{
			Reason:    "user not found",
			ErrorCode: generated.UserNotFound,
		}
	}
	return u.user, nil
}

func hashString(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}
