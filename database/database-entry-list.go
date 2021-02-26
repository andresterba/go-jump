package database

type entryList []*Entry

func (el entryList) Len() int {
	return len(el)
}
func (el entryList) Swap(i, j int) {
	el[i], el[j] = el[j], el[i]
}
func (el entryList) Less(i, j int) bool {
	return el[i].Counter > el[j].Counter
}
