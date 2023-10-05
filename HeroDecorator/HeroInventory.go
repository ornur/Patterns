package main

type IHeroInventory interface {
	getNetWorth() int
}
type HeroInventory struct {
	hero     IHeroInventory
	netWorth int
}

func (i *HeroInventory) getNetWorth() int {
	return i.netWorth
}
