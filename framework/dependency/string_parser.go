package dependency

import (
	"fmt"
	"unicode"
)

type pair struct {
	s rune
	e rune
}

type stringParser struct {
	blocks []string
}

func (o *stringParser) parse(str string, by rune) error {
	o.blocks = make([]string, 0)
	b := []rune(str)

	for i := 0; i < len(b); i++ {
		if unicode.IsSpace(b[i]) {
			continue
		}

		if b[i] == by {
			continue
		}

		// letters
		next := o.completeBlock(b, i, by)
		if next == -1 {
			return fmt.Errorf("parse error : paren are not matched for letters")
		}
		o.blocks = append(o.blocks, string(b[i:next]))
		i = next
	}

	return nil
}

func (o *stringParser) skip(b []rune, i int) int {
	i++
	for ; i < len(b); i++ {
		if b[i] == '"' {
			return i
		}
	}
	return -1
}

func (o *stringParser) completeBlock(b []rune, i int, by rune) int {
	quateIn := false
	for ; i < len(b); i++ {
		if unicode.IsSpace(b[i]) {
			continue
		}

		if b[i] == '"' {
			quateIn = !quateIn
			if !quateIn {
				continue
			}
		}

		if quateIn {
			continue
		}

		if b[i] == by {
			return i
		}

		p := isParenStarter(b[i])
		if p != nil {
			next := o.completeBlock(b, i+1, p.e)
			if next == -1 {
				return -1
			}
			i = next
		}
	}
	return len(b)
}

var pairs = []pair{
	{'"', '"'},
	{'[', ']'},
	{'{', '}'},
}

func isParenStarter(v rune) *pair {
	for _, p := range pairs {
		if p.s == v {
			return &p
		}
	}
	return nil
}
