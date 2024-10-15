package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"strings"
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
	user   *User
	config UmaConfig
}

func NewUserServiceFromEnv(config UmaConfig) *UserServiceFromEnv {
	return &UserServiceFromEnv{
		user: &User{
			ID:             config.UserID,
			Name:           config.Username,
			HashedPassword: config.HashedUserPassword,
		},
		config: config,
	}
}

func (u *UserServiceFromEnv) GetUser(id string) (*User, error) {
	if u.user == nil || u.user.ID != id {
		return nil, fmt.Errorf("user not found")
	}
	return u.user, nil
}

func (u *UserServiceFromEnv) GetUserFromContext(context *gin.Context) (*User, error) {
	if u.user == nil {
		return nil, fmt.Errorf("user not found")
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
			return nil, fmt.Errorf("error decoding credentials: %v", err)
		}
		credentials := strings.Split(string(decodedCredentials), ":")
		if len(credentials) != 2 {
			return nil, fmt.Errorf("invalid credentials")
		}
		password := credentials[1]
		hashedPassword := hashString(password)
		if credentials[0] == u.user.Name && hashedPassword == u.user.HashedPassword {
			return u.user, nil
		}
	}

	isUmaNwcReq := strings.HasPrefix(context.Request.RequestURI, "/umanwc/")
	bearerToken := "Bearer "
	if isUmaNwcReq && strings.HasPrefix(authHeader, bearerToken) {
		bearerJwt := strings.TrimPrefix(authHeader, bearerToken)
		claims, err := ParseJwt(bearerJwt, &u.config)
		if err != nil {
			return nil, fmt.Errorf("error parsing jwt: %v", err)
		}
		if claims.Subject == u.user.ID {
			return u.user, nil
		}
	}

	return nil, fmt.Errorf("user not found")
}

func (u *UserServiceFromEnv) GetUserByUmaAddress(umaAddress string, config UmaConfig, context *gin.Context) (*User, error) {
	if u.user == nil || u.user.GetUmaAddress(&config, context) != umaAddress {
		return nil, fmt.Errorf("user not found")
	}
	return u.user, nil
}

func hashString(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}
