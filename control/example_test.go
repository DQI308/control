package control

import (
	"fmt"
	"github.com/DQI308/clonebank/pkg/types"
)

func ExampleChangeActiveCard() {
	card := types.Card{}
	card.Activity = true
	ChangeActiveCard(&card)
	fmt.Println(card.Activity)
	//Output:
	//false

}