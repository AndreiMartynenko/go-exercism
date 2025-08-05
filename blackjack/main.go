package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "king", "queen", "jack":
		return 10
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
	default:
		return 0
	}
	panic("Please implement the ParseCard function")
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
// func FirstTurn(card1, card2, dealerCard string) string {

// 	panic("Please implement the FirstTurn function")
// }

func FirstTurn(card1, card2, dealerCard string) string {
	// Assign values to each card directly
	getValue := func(card string) int {
		switch card {
		case "ace":
			return 11
		case "king", "queen", "jack", "ten":
			return 10
		case "nine":
			return 9
		case "eight":
			return 8
		case "seven":
			return 7
		case "six":
			return 6
		case "five":
			return 5
		case "four":
			return 4
		case "three":
			return 3
		case "two":
			return 2
		default:
			return 0
		}
	}

	val1 := getValue(card1)
	val2 := getValue(card2)
	dealerVal := getValue(dealerCard)
	hand := val1 + val2

	// Pair of aces
	if card1 == "ace" && card2 == "ace" {
		return "P"
	}

	// Blackjack (score 21)
	if hand == 21 {
		if dealerVal < 10 {
			return "W"
		}
		return "S"
	}

	// General strategy
	if hand >= 17 {
		return "S"
	}
	if hand >= 12 && dealerVal < 7 {
		return "S"
	}
	return "H"
}
