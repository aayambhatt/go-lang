package main

import "fmt"

func main() {

	const storeName string = "THE FOODIES"
	var itemOne string = "item one"
	var itemTwo string = "item two"
	var itemThree string = "item three"

	var priceOne float64 = 12.324
	var priceTwo float64 = 54.253
	var priceThree float64 = 32.425

	receiptLineOne := fmt.Sprintf("%s.......... $%.2f\n", itemOne, priceOne)
	receiptLineTwo := fmt.Sprintf("%s.......... $%.2f\n", itemTwo, priceTwo)
	receiptLineThree := fmt.Sprintf("%s........ $%.2f\n", itemThree, priceThree)
	totalLine := fmt.Sprintf("TOTAL............. $%.2f\n", priceOne+priceTwo+priceThree)

	fmt.Print("==== " + storeName + " ====\n")
	fmt.Print(receiptLineOne)
	fmt.Print(receiptLineTwo)
	fmt.Print(receiptLineThree)

	fmt.Print("_______________________\n")
	fmt.Print(totalLine)
}
