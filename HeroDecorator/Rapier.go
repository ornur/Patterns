package main

type Rapier struct {
	hero IHeroInventory
}

func (i *Rapier) getNetWorth() int {
	return i.hero.getNetWorth() + 5600
}
