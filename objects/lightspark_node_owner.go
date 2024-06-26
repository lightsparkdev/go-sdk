// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

// LightsparkNodeOwner This is an object representing the owner of a LightsparkNode.
type LightsparkNodeOwner interface {
	Entity
}

func LightsparkNodeOwnerUnmarshal(data map[string]interface{}) (LightsparkNodeOwner, error) {
	if data == nil {
		return nil, nil
	}

	dataJSON, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	switch data["__typename"].(string) {
	case "Account":
		var account Account
		if err := json.Unmarshal(dataJSON, &account); err != nil {
			return nil, err
		}
		return account, nil
	case "Wallet":
		var wallet Wallet
		if err := json.Unmarshal(dataJSON, &wallet); err != nil {
			return nil, err
		}
		return wallet, nil

	default:
		return nil, fmt.Errorf("unknown LightsparkNodeOwner type: %s", data["__typename"])
	}
}
