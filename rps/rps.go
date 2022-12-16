package rps

import (
	"math/rand"
	"time"
)

const (
	ROCK         = 0
	PAPER        = 1
	SCISSORS     = 2
)

type Round struct {
	Feedback string `json:"feedback"`
	ComputerChoice string `json:"computer_choice"`
	RoundResult string `json:"round_result"`
}


func PlayRound(playerValue int) Round {
	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(3)
	computerChoice := ""
	roundResult := ""
	feedback := ""
	winFeedback := [3]string{
		"You're a winner!",
		"OKAY!",
		"Lucky again!",
	}
	loseFeedback := [3]string{
		"Better luck next time",
		"Unlucky this time",
		"Sorry, you lose this time",
	}
	drawFeedback := [3]string{
		"Free go!",
		"Push",
		"Try again",
	}
	
	switch computerValue {
	case ROCK:
		// *** changed here
		computerChoice = "Computer chose ROCK"
		break
	case PAPER:
		// *** changed here
		computerChoice = "Computer chose PAPER"
		break
	case SCISSORS:
		// *** changed here
		computerChoice = "Computer chose SCISSORS"
		break
	default:
	}

	if playerValue == computerValue {
		roundResult = "It's a draw"
		feedback = drawFeedback[rand.Intn(len(drawFeedback))]
	} else if playerValue == (computerValue+1)%3 {
		roundResult = "Player wins"
		feedback = winFeedback[rand.Intn(len(winFeedback))]
	} else {
		roundResult = "Computer wins"
		feedback = loseFeedback[rand.Intn(len(loseFeedback))]
	}


	var result Round
	result.Feedback = feedback
	result.ComputerChoice = computerChoice
	result.RoundResult = roundResult

	return result
}
