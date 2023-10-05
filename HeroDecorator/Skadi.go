package main

type Skadi struct {
	hero IHeroInventory
}

func (i *Skadi) getNetWorth() int {
	return i.hero.getNetWorth() + 5300
}
