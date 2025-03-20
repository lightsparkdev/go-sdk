package main

import (
	"github.com/lightsparkdev/go-sdk/objects"
	"github.com/lightsparkdev/go-sdk/services"
	"github.com/uma-universal-money-address/uma-go-sdk/uma/errors"
	"github.com/uma-universal-money-address/uma-go-sdk/uma/generated"
)

func GetNode(client *services.LightsparkClient, nodeId string) (*objects.LightsparkNode, error) {
	entity, err := client.GetEntity(nodeId)
	if err != nil {
		return nil, err
	}

	castNode, didCast := (*entity).(objects.LightsparkNode)
	if !didCast {
		return nil, &errors.UmaError{
			Reason:    "failed to cast entity to LightsparkNode",
			ErrorCode: generated.InternalError,
		}
	}

	return &castNode, nil
}
