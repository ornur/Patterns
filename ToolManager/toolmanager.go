package main

func main() {
	pickaxe := &Pickaxe{}
	shovel := &Shovel{}
	axe := &Axe{}

	workPlace := &WorkPlace{
		blocks: []string{},
	}
	workPlace.addBlock("stone")
	workPlace.addBlock("dirt")
	workPlace.addBlock("wood")
	workPlace.addBlock("sand")
	workPlace.addBlock("stone")
	workPlace.addBlock("wood")

	workPlace.setTool(pickaxe)
	workPlace.use()
	workPlace.use()
	workPlace.setTool(shovel)
	workPlace.use()
	workPlace.use()
	workPlace.setTool(axe)
	workPlace.use()
	workPlace.setTool(shovel)
	workPlace.use()
	workPlace.setTool(pickaxe)
	workPlace.use()
}
