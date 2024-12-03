package resource

import (
	"strings"
)

type Parser []byte

func (p Parser) Index(i int) byte {
	if i < 0 || i >= len(p) {
		return 0
	}

	return p[i]
}

func (p Parser) ParseLinks(fn func(link string)) {
	for i := range p {
		// find the start of the href attribute
		if p[i] == '<' && p.Index(i+1) == 'a' && p.Index(i+2) == ' ' {
			i += 3
			// find the start of the href attribute value
			for ; i < len(p); i++ {
				if p[i] == '>' {
					break
				}

				if p.Index(i) == 'h' && p.Index(i+1) == 'r' && p.Index(i+2) == 'e' && p.Index(i+3) == 'f' && p.Index(i+4) == '=' {
					i += 5

					// find the start of the href attribute value
					for ; i < len(p); i++ {
						if p.Index(i) == '"' {
							i++

							break
						}
					}

					start := i
					for ; i < len(p); i++ {
						if p.Index(i) == '"' {
							fn(strings.TrimSpace(string(p[start:i])))

							break
						}
					}

					break
				}
			}
		}
	}
}
