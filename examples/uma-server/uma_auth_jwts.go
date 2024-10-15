package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type UmaAuthJwt struct {
	Address string `json:"address"`
	jwt.RegisteredClaims
}

func NewUmaAuthJwtForUser(user *User, config *UmaConfig, context *gin.Context, expirySecs int) UmaAuthJwt {
	return UmaAuthJwt{
		Address: user.GetUmaAddress(config, context),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.OwnVaspDomain,
			Subject:   user.ID,
			Audience:  []string{config.OwnVaspDomain, *config.NwcDomain},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(expirySecs))),
		},
	}
}

func (u UmaAuthJwt) Sign(config *UmaConfig) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, u)
	ecPrivkey, err := jwt.ParseECPrivateKeyFromPEM([]byte(config.NwcJwtPrivKey))
	if err != nil {
		return "", err
	}
	return token.SignedString(ecPrivkey)
}

func ParseJwt(token string, config *UmaConfig) (UmaAuthJwt, error) {
	var claims UmaAuthJwt
	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return jwt.ParseECPublicKeyFromPEM([]byte(config.NwcJwtPubKey))
	})

	return claims, err
}
