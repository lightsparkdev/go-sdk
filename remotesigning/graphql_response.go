package remotesigning

type GraphQLResponse struct {
	Query       string
	Variables   map[string]interface{}
	OutputField string
}
