package structs

import (
	"fmt"
	"strings"
)

type Hand struct {
	Role   Role
	Cards  []Card
	Points int
}

type Role string

const (
	Player Role = "player"
	Dealer Role = "dealer"
)

func (h *Hand) Print(allHand bool) string {
	var handList []string
	if allHand {
		for _, card := range h.Cards {
			handList = append(handList, fmt.Sprintf("%s of %s", card.Name, card.Suit))
		}
	} else {
		// For visible hand only
		for _, card := range h.Cards[1:len(h.Cards)] {
			handList = append(handList, fmt.Sprintf("%s of %s", card.Name, card.Suit))
		}
	}
	return strings.Join(handList, " - ")
}

func (h *Hand) SumPoints() {
	var totalPoints int
	if h.Role == Player {
		for _, card := range h.Cards {
			var aceChoice string
			if card.Name == "Ace" {
				fmt.Println()
				for aceChoice != "11" && aceChoice != "1" {
					fmt.Print("Choose value for Ace of ", card.Suit, " [1 or 11]: ")
					fmt.Scan(&aceChoice)
					aceChoice = strings.TrimSpace(aceChoice)
					aceChoice = strings.ToLower(aceChoice)
				}
				if aceChoice == "11" {
					fmt.Println("Ace of", card.Suit, "as 11")
				} else if aceChoice == "1" {
					card.Value = 1
					fmt.Println("Ace of", card.Suit, "as 1")
				}
			}
			totalPoints += card.Value
		}
	} else if h.Role == Dealer {
		for _, card := range h.Cards {
			if card.Name == "Ace" {
				if totalPoints+card.Value > 21 {
					card.Value = 1
				}
			}
			totalPoints += card.Value
		}
	}
	h.Points = totalPoints
}
