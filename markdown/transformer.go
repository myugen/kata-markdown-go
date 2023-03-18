package markdown

import (
	"regexp"
	"strings"
)

type Transformer struct {
	text    Text
	links   []Link
	anchors Anchors
}

func (t Transformer) Text() Text {
	return t.text
}

func (t Transformer) Links() []Link {
	return t.links
}

func (t Transformer) Anchors() Anchors {
	return t.anchors
}

func (t Transformer) Transform() Text {
	transformedText := t.replaceAllLinksForAnchors()
	footnotes := t.buildFootnotes()
	return transformedText + "\n\n" + footnotes
}

func (t Transformer) replaceAllLinksForAnchors() Text {
	transformedText := t.text
	for _, link := range t.links {
		transformedText = strings.ReplaceAll(transformedText, link.AsText(), t.anchors.AnchorFor(link))
	}
	return transformedText
}

func (t Transformer) buildFootnotes() string {
	return t.anchors.ToFootnote()
}

func NewTransformer(text Text) *Transformer {
	regex := regexp.MustCompile(markdownLinkRegex)
	links := make([]Link, 0)
	anchors := NewAnchors()
	for _, foundGroup := range regex.FindAllStringSubmatch(text, -1) {
		label := strings.Replace(strings.Replace(foundGroup[1], "[", "", -1), "]", "", -1)
		url := strings.Replace(strings.Replace(foundGroup[2], "(", "", -1), ")", "", -1)
		links = append(links, NewLink(label, url))
		if !anchors.Exists(url) {
			anchors.Add(url)
		}
	}
	return &Transformer{text: text, links: links, anchors: anchors}
}
