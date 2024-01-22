package main

import (
	"errors"
	"github.com/lightsparkdev/go-sdk/objects"
	"github.com/lightsparkdev/go-sdk/services"
)

func GetNode(client *services.LightsparkClient, nodeId string) (*objects.LightsparkNode, error) {
	entity, err := client.GetEntity(nodeId)
	if err != nil {
		return nil, err
	}

	castNode, didCast := (*entity).(objects.LightsparkNode)
	if !didCast {
		return nil, errors.New("failed to cast entity to LightsparkNode")
	}

	return &castNode, nil
}
