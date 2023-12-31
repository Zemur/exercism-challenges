package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
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
        case "one":
        	return 1
        default:
    		return 0
    }
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    var cardSum int = ParseCard(card1) + ParseCard(card2)
    switch {
        case card1 == "ace" && card2 == "ace":
        	return "P"
        case cardSum == 21 && ParseCard(dealerCard) < 10:
        	return "W"
        case cardSum == 21 && ParseCard(dealerCard) >= 10:
        	return "S"
        case cardSum >= 17 && cardSum <= 20:
        	return "S"
        case cardSum >= 12 && cardSum <= 16 && ParseCard(dealerCard) < 7:
        	return "S"
        case cardSum >= 12 && cardSum <= 16 && ParseCard(dealerCard) >= 7:
        	return "H"
        default:
        	return "H"
    }
}
