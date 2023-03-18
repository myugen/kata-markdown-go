package markdown_test

import (
	"testing"

	"github.com/AgileCraftsmanshipCanarias/kata-setup-go/markdown"
	"github.com/stretchr/testify/assert"
)

func TestAnchors_AnchorFor(t *testing.T) {
	type fields struct {
		urls []markdown.URL
	}
	type args struct {
		link markdown.Link
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "should retrieve formatted text with correspondent anchor for provided link",
			fields: fields{[]markdown.URL{"https://www.codigosostenible.es"}},
			args:   args{markdown.NewLink("El libro código sostenible", "https://www.codigosostenible.es")},
			want:   "El libro código sostenible[^anchor1]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			anchors := markdown.NewAnchors(tt.fields.urls...)
			got := anchors.AnchorFor(tt.args.link)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAnchors_ToFootnote(t *testing.T) {
	type fields struct {
		urls []markdown.URL
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "should map all anchors as footnote text",
			fields: fields{[]markdown.URL{"https://www.codigosostenible.com", "https://wikipedia.org"}},
			want:   "[^anchor1]: https://www.codigosostenible.com\n[^anchor2]: https://wikipedia.org",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			anchors := markdown.NewAnchors(tt.fields.urls...)
			got := anchors.ToFootnote()
			assert.Equal(t, tt.want, got)
		})
	}
}
