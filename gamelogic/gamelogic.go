package gamelogic

import (
	"fmt"
	"math/rand"
	"time"
)

// GameCard defines a card schema
type GameCard struct {
	Type string
}

// GamePlay defines a gameplay by a player
type GamePlay struct {
	Card   GameCard
	Player string
}

// Trophy defines the game results
type Trophy struct {
	Winner     GamePlay
	Looser     GamePlay
	ResultType string
}

// GameLogic holds the game logic function
type GameLogic struct{}

// Result returns result
func (g *GameLogic) Result(winner GamePlay, looser GamePlay, desc string) Trophy {
	var result Trophy

	result.Winner = winner
	result.Looser = looser
	result.ResultType = desc

	if desc == "DRAW" {
		fmt.Println(
			result.Winner.Player +
				" Draws with " +
				result.Looser.Player +
				"\n" +
				result.Winner.Card.Type +
				" draws.",
		)
	} else {
		fmt.Println(
			"Winner: " +
				result.Winner.Player +
				"\n" + "Looser: " +
				result.Looser.Player + "\n" +
				result.Winner.Card.Type +
				" beats " +
				result.Looser.Card.Type,
		)
	}

	return result
}

// Play handles the game play
func (g *GameLogic) Play(playerOneCard GamePlay, playerTwoCard GamePlay) {
	/**
	Scisoor beats Paper
	Paper beats Rock
	Rock beats Sciscor
	*/

	if playerOneCard.Card.Type == "SCISSOR" && playerTwoCard.Card.Type == "ROCK" {
		g.Result(playerTwoCard, playerOneCard, "COMPLETE")
	} else if playerOneCard.Card.Type == "SCISSOR" && playerTwoCard.Card.Type == "PAPER" {
		g.Result(playerOneCard, playerTwoCard, "COMPLETE")
	} else if playerOneCard.Card.Type == "SCISSOR" && playerTwoCard.Card.Type == "SCISSOR" {
		g.Result(playerOneCard, playerTwoCard, "DRAW")
	} else if playerOneCard.Card.Type == "ROCK" && playerTwoCard.Card.Type == "PAPER" {
		g.Result(playerTwoCard, playerOneCard, "COMPLETE")
	} else if playerOneCard.Card.Type == "ROCK" && playerTwoCard.Card.Type == "SCISSOR" {
		g.Result(playerOneCard, playerTwoCard, "COMPLETE")
	} else if playerOneCard.Card.Type == "ROCK" && playerTwoCard.Card.Type == "ROCK" {
		g.Result(playerOneCard, playerTwoCard, "DRAW")
	} else if playerOneCard.Card.Type == "PAPER" && playerTwoCard.Card.Type == "SCISSOR" {
		g.Result(playerTwoCard, playerOneCard, "COMPLETE")
	} else if playerOneCard.Card.Type == "PAPER" && playerTwoCard.Card.Type == "ROCK" {
		g.Result(playerOneCard, playerTwoCard, "COMPLETE")
	} else if playerOneCard.Card.Type == "PAPER" && playerTwoCard.Card.Type == "PAPER" {
		g.Result(playerOneCard, playerTwoCard, "DRAW")
	}
}

// Computer play handles computer gameplay
func (g *GameLogic) Computer() GamePlay {
	var cards []string = []string{"ROCK", "PAPER", "SCISSOR"}

	rand.Seed(time.Now().Unix())

	cardType := cards[rand.Intn(len(cards))]

	var gameplay GamePlay = GamePlay{
		Card: GameCard{
			Type: cardType,
		},
		Player: "Computer",
	}

	return gameplay
}
