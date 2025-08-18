package main

import "fmt"

// import "fmt"

// func main() {
// 	var x int = 0
// 	switch  {
// 	case x < 0:
// 		fmt.Println("Negative")
// 	case x > 0:
// 		fmt.Println("Positive")

// 	default :
// 	 fmt.Println("Good Morning")
// 	}
// }

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    switch card {
        case "ace":
         return 11
		case "king":
			return 10
    }
	return 0
} 
func main() {
	value:= ParseCard("king")
	fmt.Println("card",value)
}
