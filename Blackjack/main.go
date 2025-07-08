package main

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
        case "ace": return 11
        case "two": return 2
        case "three": return 3
        case "four": return 4
        case "five": return 5
        case "six": return 6
        case "seven": return 7
        case "eight": return 8
        case "nine": return 9
        case "jack", "queen", "king", "ten": return 10
        default: return 0
    }
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    parsedCard1 := ParseCard(card1)
    parsedCard2 := ParseCard(card2)
    parsedDealerCard := ParseCard(dealerCard)

    sumCards := parsedCard1 + parsedCard2
    switch {
        case parsedCard1 == 11 && parsedCard2 == 11:
        	return "P"
        case (sumCards == 21):
        	if (parsedDealerCard == 11 || parsedDealerCard == 10) {
                return "S"
            }
        	return "W"
        case sumCards >= 17 && sumCards <= 20:
        	return "S"
        case sumCards >= 12 && sumCards <= 16:
        	if parsedDealerCard >= 7 {
                return "H"
            }
        	return "S"
        default:
        	return "H"
    }

    
}
