package main

import (
	"errors"
	"github.com/lightsparkdev/go-sdk/objects"
	"github.com/lightsparkdev/go-sdk/services"
	"time"
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

func WaitForPaymentCompletion(client *services.LightsparkClient, payment *objects.OutgoingPayment) (*objects.OutgoingPayment, error) {
	attemptLimit := 200
	for payment.Status != objects.TransactionStatusSuccess && payment.Status != objects.TransactionStatusFailed {
		if attemptLimit == 0 {
			return nil, errors.New("payment timed out")
		}
		attemptLimit--
		time.Sleep(100 * time.Millisecond)

		entity, err := client.GetEntity(payment.Id)
		if err != nil {
			return nil, err
		}
		castPayment, didCast := (*entity).(objects.OutgoingPayment)
		if !didCast {
			return nil, errors.New("failed to cast payment to OutgoingPayment")
		}
		payment = &castPayment
	}
	return payment, nil
}
