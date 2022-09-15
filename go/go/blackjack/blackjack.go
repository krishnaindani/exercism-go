package blackjack

const (
	Split            = "P"
	Stand            = "S"
	Hit              = "H"
	AutomaticallyWin = "W"
)

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "ace":
		return 11
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {

	dealerCardValue := ParseCard(dealerCard)
	cardSum := ParseCard(card1) + ParseCard(card2)

	switch {
	case card1 == "ace" && card2 == "ace":
		return Split
	case cardSum == 21:
		if dealerCardValue != 11 && dealerCardValue != 10 {
			return AutomaticallyWin
		}
		return Stand
	case cardSum >= 17 && cardSum <= 20:
		return Stand
	case cardSum >= 12 && cardSum <= 16:
		if dealerCardValue >= 7 {
			return Hit
		}
		return Stand
	case cardSum <= 11:
		return Hit
	}
	return ""
}
