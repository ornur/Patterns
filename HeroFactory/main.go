package main

import "fmt"

func main() {
	rubick, _ := getHero("Rubick")
	dazzle, _ := getHero("Dazzle")
	drow, _ := getHero("Drow Ranger")
	tiny, _ := getHero("Tiny")

	printHero(rubick)
	printHero(dazzle)
	printHero(drow)
	printHero(tiny)
}

func printHero(hero IHero) {
	fmt.Println(hero.getName() + " is a " + hero.getType() + " hero.")
}
