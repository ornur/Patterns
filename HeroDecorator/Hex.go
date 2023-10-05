package main

type Hex struct {
	hero IHeroInventory
}

func (i *Hex) getNetWorth() int {
	return i.hero.getNetWorth() + 5550
}
