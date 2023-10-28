package main

func main() {
	armlet := &Armlet{}
	radiance := &Radiance{}

	oncommand := &OnCommand{
		toggleItem: armlet,
	}
	offcommand := &OffCommand{
		toggleItem: armlet,
	}
	onKeyPress := &KeyPress{
		command: oncommand,
	}
	onKeyPress.press()

	offKeyPress := &KeyPress{
		command: offcommand,
	}
	offKeyPress.press()

	oncommand = &OnCommand{
		toggleItem: radiance,
	}
	offcommand = &OffCommand{
		toggleItem: radiance,
	}
	onKeyPress.press()
	offKeyPress.press()

}
