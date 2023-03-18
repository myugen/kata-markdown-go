package markdown

import (
	"fmt"
	"strings"
)

type Anchors struct {
	urls []URL
}

func (a *Anchors) Add(url URL) {
	a.urls = append(a.urls, url)
}

func (a *Anchors) Exists(url URL) bool {
	for _, u := range a.urls {
		if u == url {
			return true
		}
	}
	return false
}

func (a *Anchors) AnchorFor(link Link) Text {
	index := 0
	for i, url := range a.urls {
		if url == link.Url() {
			index = i
		}
	}
	return fmt.Sprintf("%s[^anchor%d]", link.Label(), index+1)
}

func (a *Anchors) ToFootnote() Text {
	formattedAnchors := make([]string, 0)
	for i, url := range a.urls {
		formattedAnchor := fmt.Sprintf("[^anchor%d]: %s", i+1, url)
		formattedAnchors = append(formattedAnchors, formattedAnchor)
	}
	return strings.Join(formattedAnchors, "\n")
}

func NewAnchors(urls ...URL) Anchors {
	return Anchors{urls: urls}
}
