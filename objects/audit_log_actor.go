// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
	"fmt"
)

// AuditLogActor Audit log actor who called the GraphQL mutation
type AuditLogActor interface {
	Entity

	// GetTypename The typename of the object
	GetTypename() string
}

func AuditLogActorUnmarshal(data map[string]interface{}) (AuditLogActor, error) {
	if data == nil {
		return nil, nil
	}

	dataJSON, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	switch data["__typename"].(string) {
	case "ApiToken":
		var apiToken ApiToken
		if err := json.Unmarshal(dataJSON, &apiToken); err != nil {
			return nil, err
		}
		return apiToken, nil

	default:
		return nil, fmt.Errorf("unknown AuditLogActor type: %s", data["__typename"])
	}
}
