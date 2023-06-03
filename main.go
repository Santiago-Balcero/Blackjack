package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"blackjack/constants"
	"blackjack/structs"
)

// Gets a new card from Deck and deletes it from Deck
func selectCard() structs.Card {
	index := rand.Intn(len(constants.Deck))
	card := constants.Deck[index]
	constants.Deck = append(constants.Deck[:index], constants.Deck[index+1:]...)
	return card
}

// Allows a player to play
func playPlayerHand(hand structs.Hand) structs.Hand {
	for {
		var choice string
		card := selectCard()
		time.Sleep(2 * time.Second)
		fmt.Println("You got:", card.Name, "of", card.Suit)
		hand.Cards = append(hand.Cards, card)
		fmt.Println("Your hand:", hand.Print(true))
		hand.SumPoints("player")
		if hand.Points >= 21 {
			break
		}
		fmt.Println()
		for choice != "y" && choice != "n" {
			fmt.Print("Another card? [y/n]: ")
			fmt.Scan(&choice)
			choice = strings.TrimSpace(choice)
			choice = strings.ToLower(choice)
		}
		if choice == "y" {
			fmt.Println()
			continue
		} else {
			break
		}
	}
	return hand
}

// Allows dealer to play
func playDealerHand(hand structs.Hand) structs.Hand {
	for {
		card := selectCard()
		time.Sleep(2 * time.Second)
		fmt.Println("\nDealer got:", card.Name, "of", card.Suit)
		hand.Cards = append(hand.Cards, card)
		fmt.Println("Dealer's visible hand:", hand.Print(false))
		hand.SumPoints("dealer")
		if hand.Points >= 17 {
			break
		}
	}
	return hand
}

func checkWinner(playerHand, dealerHand structs.Hand) structs.Hand {
	fmt.Println("\nChecking hands for winner...")
	time.Sleep(2 * time.Second)
	playerDiff := 21 - playerHand.Points
	dealerDiff := 21 - dealerHand.Points
	if (playerDiff == 0 && dealerDiff != 0) || (playerDiff > 0 && dealerDiff < 0) {
		fmt.Println("\nYOU WIN!")
		return playerHand
	} else if (dealerDiff == 0 && playerDiff != 0) || (dealerDiff > 0 && playerDiff < 0) {
		fmt.Println("\nYOU LOSE!")
		return dealerHand
	} else {
		return checkCards(playerHand, dealerHand)
	}
}

func checkCards(playerHand, dealerHand structs.Hand) structs.Hand {
	// Returns an empty hand with 0 points if there was a tie
	if len(playerHand.Cards) == 2 && playerHand.Points == 21 && len(dealerHand.Cards) > 2 {
		// Player Blackjack
		fmt.Println("\nYOU WIN!")
		return playerHand
	} else if len(dealerHand.Cards) == 2 && dealerHand.Points == 21 && len(playerHand.Cards) > 2 {
		// Dealer Blackjack
		fmt.Println("\nYOU LOSE!")
		return dealerHand
	} else {
		fmt.Println("\nTIED GAME!")
		return structs.Hand{}
	}
}

func main() {
	fmt.Println("\nBLACKJACK GAME")
	deck := make([]structs.Card, len(constants.Deck))
	copy(deck, constants.Deck)
	dealerCard1 := selectCard()
	dealerCard2 := selectCard()
	dealerHand := structs.Hand{
		Cards:  []structs.Card{dealerCard1, dealerCard2},
		Points: dealerCard1.Value + dealerCard2.Value,
	}
	time.Sleep(2 * time.Second)
	fmt.Println("\nVisible card of dealer is:", dealerCard2.Name, "of", dealerCard2.Suit)
	// fmt.Println(dealerHand)
	playerCard := selectCard()
	time.Sleep(2 * time.Second)
	fmt.Println("\nYou got:", playerCard.Name, "of", playerCard.Suit)
	playerHand := structs.Hand{
		Cards:  []structs.Card{playerCard},
		Points: playerCard.Value,
	}
	playerHand = playPlayerHand(playerHand)
	// Case player points > 21
	if playerHand.Points <= 21 {
		dealerHand = playDealerHand(dealerHand)
	}
	winnerHand := checkWinner(playerHand, dealerHand)
	if winnerHand.Points == 0 {
		fmt.Println("No one wins")
	} else {
		fmt.Println("Winner points:", winnerHand.Points)
		fmt.Println("Winner hand:", winnerHand.Print(true))
		if winnerHand.Points == 21 && len(winnerHand.Cards) == 2 {
			fmt.Println("BLACKJACK!")
		}
	}
	fmt.Println("")
}
