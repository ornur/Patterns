package main

type KeyPress struct {
	command Command
}

func (press *KeyPress) press() {
	press.command.execute()
}

type OnCommand struct {
	toggleItem ToggleItem
}

func (on *OnCommand) execute() {
	on.toggleItem.on()
}

type OffCommand struct {
	toggleItem ToggleItem
}

func (off *OffCommand) execute() {
	off.toggleItem.off()
}
