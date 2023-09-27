package main

import "fmt"

type Pickaxe struct {
}

func (p *Pickaxe) use(w *WorkPlace) {
	if w.blocks[0] == "stone" {
		fmt.Printf("You have used a pickaxe to mine %s\n", w.blocks[0])
		w.blocks = append(w.blocks[:0], w.blocks[0+1:]...)
	} else {
		fmt.Printf("You can not use a pickaxe to mine %s\n", w.blocks[0])
	}
}

type Shovel struct {
}

func (s *Shovel) use(w *WorkPlace) {
	if w.blocks[0] == "dirt" || w.blocks[0] == "sand" {
		fmt.Printf("You have used a shovel to dig %s\n", w.blocks[0])
		w.blocks = append(w.blocks[:0], w.blocks[0+1:]...)
	} else {
		fmt.Printf("You can not use a shovel to dig %s\n", w.blocks[0])
	}
}

type Axe struct {
}

func (a *Axe) use(w *WorkPlace) {
	if w.blocks[0] == "wood" {
		fmt.Printf("You have used an axe to chop %s\n", w.blocks[0])
		w.blocks = append(w.blocks[:0], w.blocks[0+1:]...)
	} else {
		fmt.Printf("You can not use an axe to chop %s\n", w.blocks[0])
	}
}
