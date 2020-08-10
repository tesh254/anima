package main

import (
	"github.com/tesh254/anima/gamelogic"
)

func main() {
	var game gamelogic.GameLogic

	var playerOne gamelogic.GamePlay = gamelogic.GamePlay{
		Card: gamelogic.GameCard{
			Type: "SCISSOR",
		},
		Player: "Erick",
	}

	game.Play(playerOne, game.Computer())
}
