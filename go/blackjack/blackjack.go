package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
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
	case "joker":
		return 0
	default:
		return 10
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	total := ParseCard(card1) + ParseCard(card2)
	dealerTotal := ParseCard(dealerCard)

	switch {
	case total == 22:
		return "P"
	case total == 21 && dealerTotal < 10:
		return "W"
	case total >= 17 && total <= 21:
		return "S"
	case total >= 12 && total <= 16 && dealerTotal >= 7:
		return "H"
	case total >= 12 && total <= 16 && dealerTotal < 7:
		return "S"
	case total <= 11:
		return "H"
	default:
		return "A"
	}
}
