package main

type Tarassque struct {
	hero IHeroInventory
}

func (i *Tarassque) getNetWorth() int {
	return i.hero.getNetWorth() + 5000
}
