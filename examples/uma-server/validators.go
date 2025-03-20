package main

import (
	"github.com/uma-universal-money-address/uma-go-sdk/uma/errors"
	"github.com/uma-universal-money-address/uma-go-sdk/uma/generated"
	umautils "github.com/uma-universal-money-address/uma-go-sdk/uma/utils"
	"regexp"
	"strings"
)

func ValidateUmaAddress(address string) error {
	addressParts := strings.Split(address, "@")
	if len(addressParts) != 2 {
		return &errors.UmaError{
			Reason:    "invalid receiver address",
			ErrorCode: generated.InvalidInput,
		}
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
		return &errors.UmaError{
			Reason:    "invalid UMA user name",
			ErrorCode: generated.InvalidInput,
		}
	}
	return nil
}

func ValidateDomain(domain string) error {
	hostWithoutPort := strings.Split(domain, ":")[0]
	isLocalDomain := umautils.IsDomainLocalhost(hostWithoutPort)
	localHostWithPortRegex := regexp.MustCompile(`^localhost(:[0-9]+)?$`)
	domainRegex := regexp.MustCompile(`^([a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62}){1}(\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*[._]?$`)
	if !domainRegex.MatchString(domain) && !localHostWithPortRegex.MatchString(domain) && !isLocalDomain {
		return &errors.UmaError{
			Reason:    "invalid VASP domain",
			ErrorCode: generated.InvalidInput,
		}
	}
	return nil
}
