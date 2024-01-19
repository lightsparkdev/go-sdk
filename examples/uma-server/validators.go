package main

import (
	"errors"
	"github.com/uma-universal-money-address/uma-go-sdk/uma"
	"regexp"
	"strings"
)

func ValidateUmaAddress(address string) error {
	addressParts := strings.Split(address, "@")
	if len(addressParts) != 2 {
		return errors.New("invalid receiver address")
	}
	receiverId := addressParts[0]
	receiverVasp := addressParts[1]
	userNameError := ValidateUserName(receiverId)
	if userNameError != nil {
		return userNameError
	}
	domainError := ValidateDomain(receiverVasp)
	if domainError != nil {
		return domainError
	}
	return nil
}

func ValidateUserName(userName string) error {
	userNameRegex := regexp.MustCompile(`^\$?[a-zA-Z0-9-._+]+$`)
	if !userNameRegex.MatchString(userName) {
		return errors.New("invalid UMA user name")
	}
	return nil
}

func ValidateDomain(domain string) error {
	hostWithoutPort := strings.Split(domain, ":")[0]
	isLocalDomain := uma.IsDomainLocalhost(hostWithoutPort)
	localHostWithPortRegex := regexp.MustCompile(`^localhost(:[0-9]+)?$`)
	domainRegex := regexp.MustCompile(`^([a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62}){1}(\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*[._]?$`)
	if !domainRegex.MatchString(domain) && !localHostWithPortRegex.MatchString(domain) && !isLocalDomain {
		return errors.New("invalid VASP domain")
	}
	return nil
}
