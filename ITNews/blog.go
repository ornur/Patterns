package main

type Blog struct {
	observersList       []Observer
	text                string
	programmingLanguage string
}

func (b *Blog) register(o Observer) {
	b.observersList = append(b.observersList, o)
}
func (b *Blog) deregister(o Observer) {
	b.observersList = removeFromSlice(b.observersList, o)
}
func (b *Blog) notifyAll() {
	for _, observer := range b.observersList {
		observer.update(b.text)
	}
}
func (b *Blog) writeNewPost(name string) {
	b.text = name
	b.notifyAll()
}
func removeFromSlice(observersList []Observer, observerToRemove Observer) []Observer {
	observerToRemoveID := observerToRemove.getID()
	for i := 0; i < len(observersList); i++ {
		if observersList[i].getID() == observerToRemoveID {
			observersList = append(observersList[:i], observersList[i+1:]...)
			break
		}
	}
	return observersList
}
