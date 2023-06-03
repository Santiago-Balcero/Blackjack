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
		fmt.Println("You got:", card.Name, "of", card.Suit)
		time.Sleep(2 * time.Second)
		hand.Cards = append(hand.Cards, card)
		fmt.Println("Your hand:", hand.Print(true))
		time.Sleep(2 * time.Second)
		hand.SumPoints()
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
		fmt.Println("\nDealer got:", card.Name, "of", card.Suit)
		time.Sleep(2 * time.Second)
		hand.Cards = append(hand.Cards, card)
		fmt.Println("Dealer's visible hand:", hand.Print(true))
		time.Sleep(2 * time.Second)
		hand.SumPoints()
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
	} else if playerDiff > 0 && dealerDiff > 0 {
		if playerDiff < dealerDiff {
			fmt.Println("\nYOU WIN!")
			return playerHand
		} else if dealerDiff < playerDiff {
			fmt.Println("\nYOU LOSE!")
			return dealerHand
		}
	}
	return checkCards(playerHand, dealerHand)
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
	time.Sleep(2 * time.Second)
	// For when different games are played in single session
	// deck := make([]structs.Card, len(constants.Deck))
	// copy(deck, constants.Deck)

	// Dealer's starting cards
	dealerCard1 := selectCard()
	fmt.Println("\nDealer got a hole card")
	time.Sleep(2 * time.Second)
	dealerCard2 := selectCard()
	dealerHand := structs.Hand{
		Role:   structs.Dealer,
		Cards:  []structs.Card{dealerCard1, dealerCard2},
		Points: dealerCard1.Value + dealerCard2.Value,
	}
	fmt.Println("Visible card of dealer is:", dealerCard2.Name, "of", dealerCard2.Suit)
	time.Sleep(2 * time.Second)

	// Player's starting cards and turn
	playerCard := selectCard()
	fmt.Println("\nYou got:", playerCard.Name, "of", playerCard.Suit)
	time.Sleep(2 * time.Second)
	playerHand := structs.Hand{
		Role:   structs.Player,
		Cards:  []structs.Card{playerCard},
		Points: playerCard.Value,
	}
	playerHand = playPlayerHand(playerHand)

	// Dealer's turn
	// Case player points > 21 it is end game
	// Case dealer points > 17 it must stand
	if playerHand.Points <= 21 && dealerHand.Points < 17 {
		dealerHand = playDealerHand(dealerHand)
	}

	// Check for winner
	winnerHand := checkWinner(playerHand, dealerHand)
	time.Sleep(1 * time.Second)
	if winnerHand.Points == 0 {
		fmt.Println("No one wins")
	} else {
		fmt.Println("Winner points:", winnerHand.Points)
		fmt.Println("Winner hand:", winnerHand.Print(true))
		if winnerHand.Points == 21 && len(winnerHand.Cards) == 2 {
			fmt.Println("BLACKJACK!")
		}
	}
	fmt.Println()
}
