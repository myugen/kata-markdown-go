package markdown_test

import (
	"testing"

	"github.com/AgileCraftsmanshipCanarias/kata-setup-go/markdown"
	"github.com/stretchr/testify/assert"
)

func TestTransformer_Transform(t *testing.T) {
	type fields struct {
		text markdown.Text
	}
	tests := []struct {
		name   string
		fields fields
		want   markdown.Text
	}{
		{
			name:   "should not transform the text when there is no any link",
			fields: fields{"foo"},
			want:   "foo",
		},
		{
			name:   "should transform the text creating anchor links in footer",
			fields: fields{"# Test file\n\nEl [libro de Código Sostenible](https://www.codigosostenible.com) es un librazo.\n¡Cómpralo!"},
			want:   "# Test file\n\nEl libro de Código Sostenible[^anchor1] es un librazo.\n¡Cómpralo!\n\n[^anchor1]: https://www.codigosostenible.com",
		},
		{
			name:   "should transform the text creating anchor for links in footer without repeating",
			fields: fields{"# Test file\n\nEl [libro de Código Sostenible](https://www.codigosostenible.com) es un librazo.\nDe verdad, [este libro](https://www.codigosostenible.com) es un librazo.\n¡Cómpralo!"},
			want:   "# Test file\n\nEl libro de Código Sostenible[^anchor1] es un librazo.\nDe verdad, este libro[^anchor1] es un librazo.\n¡Cómpralo!\n\n[^anchor1]: https://www.codigosostenible.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transformer := markdown.NewTransformer(tt.fields.text)
			got := transformer.Transform()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNewTransformer(t *testing.T) {
	type args struct {
		text markdown.Text
	}
	type want struct {
		links   []markdown.Link
		anchors markdown.Anchors
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "should create transformer with detected links and all anchors without repeating urls",
			args: args{"# Test file\n\nEl [libro de Código Sostenible](https://www.codigosostenible.com) es un librazo.\n\nEl [mismo libro](https://www.codigosostenible.com) es un librazo.\n¡Cómpralo!"},
			want: want{
				links: []markdown.Link{
					markdown.NewLink("libro de Código Sostenible", "https://www.codigosostenible.com"),
					markdown.NewLink("mismo libro", "https://www.codigosostenible.com"),
				},
				anchors: markdown.NewAnchors("https://www.codigosostenible.com"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := markdown.NewTransformer(tt.args.text)
			assert.Equal(t, tt.want.links, got.Links())
			assert.Equal(t, tt.want.anchors, got.Anchors())
		})
	}
}
