package main

type Butterfly struct {
	hero IHeroInventory
}

func (i *Butterfly) getNetWorth() int {
	return i.hero.getNetWorth() + 4975
}
