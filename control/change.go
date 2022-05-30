package control

import "github.com/DQI308/clonebank/pkg/types"

func ChangeActiveCard(card *types.Card) {
	(*card).Activity = false
}