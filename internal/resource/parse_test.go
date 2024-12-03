package resource

import (
	"net/url"
	"reflect"
	"testing"
)

func TestParser_Index(t *testing.T) {
	tests := []struct {
		name string
		p    Parser
		args []int
		want []byte
	}{
		{
			name: "index 0",
			p:    Parser("test"),
			args: []int{0, 1, 5, 10, -1, 999},
			want: []byte(`te`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := range tt.args {
				if got := tt.p.Index(tt.args[i]); got != 0 && got != tt.want[i] {
					t.Errorf("Parser.Index() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestParser_ParseLinks(t *testing.T) {
	tests := []struct {
		name     string
		p        Parser
		baseLink string
		want     []string
	}{
		{
			name:     "parse links",
			p:        Parser([]byte(`<a href="https://example.com">`)),
			baseLink: "https://example.com",
			want:     []string{"https://example.com"},
		},
		{
			name:     "parse empty links",
			p:        Parser([]byte(`<a href="">`)),
			baseLink: "https://example.com",
			want:     nil,
		},
		{
			name:     "parse multiple links",
			p:        Parser([]byte(`<a href="https://example.com">    <a href="https://test.com">`)),
			baseLink: "https://example.com",
			want:     []string{"https://example.com"},
		},
		{
			name:     "parse multiple links",
			p:        Parser([]byte(`<a href="https://start.url/abc/foo"> <a href="https://start.url/abc/foo/bar"> <a href="https://start.url/baz"> <a href="https://another.domain">`)),
			baseLink: "https://start.url/abc",
			want:     []string{"https://start.url/abc/foo", "https://start.url/abc/foo/bar"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseURLParse, err := url.Parse(tt.baseLink)
			if err != nil {
				t.Fatalf("failed to parse base link: %v", err)
			}

			var got []string
			tt.p.ParseLinks(func(link string) {
				if ok, _ := isRelatedLink(baseURLParse, link); ok {
					got = append(got, link)
				}
			})

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.ParseLinks() = %v, want %v", got, tt.want)
			}
		})
	}
}
