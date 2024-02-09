package main

import (
	"github.com/merliot/hp2430n"
	"github.com/merliot/hub"
	"github.com/merliot/ps30m"
)

func registerModels(hub *hub.Hub) {
	hub.RegisterModel("hp2430n", hp2430n.New)
	hub.RegisterModel("ps30m", ps30m.New)
}
