package models

type Dialog struct {
	title  string
	choice bool
}

func CreateDialog(title string) Dialog {
	return Dialog{
		title: title,
	}
}

func (d *Dialog) SetChoice(t bool) {
	d.choice = t
}

func (d *Dialog) GetQuestion() string {
	return d.title
}

func (d *Dialog) OnClick(checkHandler func()) {
	if d.choice {
		checkHandler()
	}
}
