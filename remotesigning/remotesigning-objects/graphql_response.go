// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

type GraphQLResponse struct {
	Query       string
	Variables   map[string]interface{}
	OutputField string
}
