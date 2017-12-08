package cmd

import (
	"github.com/kochman/runner"

	"github.com/kochman/partyplayer/config"
	"github.com/kochman/partyplayer/controller"
	"github.com/kochman/partyplayer/log"
	"github.com/kochman/partyplayer/server"
)

func Run() {
	log.Info("PartyPlayer starting...")
	log.SetLevel("debug")

	config, err := config.New()
	if err != nil {
		log.WithError(err).Error("Unable to create config.")
		return
	}
	runner := runner.New()

	controller, err := controller.New(config)
	if err != nil {
		log.WithError(err).Error("Unable to create controller.")
		return
	}

	server, err := server.New(config, controller)
	if err != nil {
		log.WithError(err).Error("Unable to create server.")
		return
	}
	runner.Add(server)

	// player, err := player.New(config)
	// if err != nil {
	// 	log.WithError(err).Error("Unable to create player.")
	// 	return
	// }
	// runner.Add(player)

	runner.Run()
}
