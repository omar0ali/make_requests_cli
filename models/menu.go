package models

var counter = -1

type MenuItem struct {
	ID    int
	Title string
}

func CreateItem(title string) MenuItem {
	counter++
	return MenuItem{
		ID:    counter,
		Title: title,
	}
}

func (item *MenuItem) OnClick(clickHandler func()) {
	clickHandler()
}
