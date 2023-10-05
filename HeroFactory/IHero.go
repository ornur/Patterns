package main

type IHero interface {
	getName() string
	setName(name string)
	getType() string
	setType(heroType string)
}
