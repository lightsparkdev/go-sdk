package main

import (
	"github.com/lightsparkdev/go-sdk/services"
)

type UmaAuthAdapter struct {
	sendingVasp *Vasp1
	config      *UmaConfig
	client      *services.LightsparkClient
}
