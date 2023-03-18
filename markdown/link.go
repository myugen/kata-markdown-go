package markdown

import "fmt"

type Link struct {
	label Label
	url   URL
}

func (l Link) Label() Label {
	return l.label
}

func (l Link) Url() URL {
	return l.url
}

func (l Link) AsText() Text {
	return fmt.Sprintf("[%s](%s)", l.label, l.url)
}

func NewLink(label Label, url URL) Link {
	return Link{label: label, url: url}
}
