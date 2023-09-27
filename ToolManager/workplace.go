package main

type WorkPlace struct {
	blocks []string
	tool   Tool
}

func (w *WorkPlace) setTool(tool Tool) {
	w.tool = tool
}
func (w *WorkPlace) use() {
	if w.tool != nil && len(w.blocks) > 0 {
		w.tool.use(w)
	}
}
func (w *WorkPlace) addBlock(block string) {
	w.blocks = append(w.blocks, block)
}
func (w *WorkPlace) getBlocks() []string {
	return w.blocks
}
