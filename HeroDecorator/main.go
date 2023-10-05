package main

import "fmt"

func main() {
	HeroInventory := &HeroInventory{}

	InventoryWithSkadi := &Skadi{HeroInventory}
	fmt.Println("Net worth with Skadi: ", InventoryWithSkadi.getNetWorth())

	InventoryWithSkadiAndButterfly := &Butterfly{InventoryWithSkadi}
	fmt.Println("Net worth with Skadi and Butterfly: ", InventoryWithSkadiAndButterfly.getNetWorth())

	AddedTarrasque := &Tarassque{InventoryWithSkadiAndButterfly}
	fmt.Println("Net worth with Skadi, Butterfly and Tarrasque: ", AddedTarrasque.getNetWorth())

	AddedRapier := &Rapier{AddedTarrasque}
	fmt.Println("Net worth with Skadi, Butterfly, Tarrasque and Rapier: ", AddedRapier.getNetWorth())

	FinalInventory := &Hex{AddedRapier}
	fmt.Println("Net worth with Skadi, Butterfly, Tarrasque, Rapier and Hex: ", FinalInventory.getNetWorth(), "\n")
	fmt.Println("Total net worth: ", FinalInventory.getNetWorth())
}
