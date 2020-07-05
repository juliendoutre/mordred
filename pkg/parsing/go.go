package parsing

import (
	"fmt"
	"strings"
)

// GoParser captures strings in golang files.
type GoParser struct {
	inChar               bool
	inSingleLineComment  bool
	inMultiLineComment   bool
	inDoubleQuotedString bool
	inBacktickedString   bool
	previousRune         rune
	strings              []string
}

// Parse returns the strings captured in a golang file contents.
func (g *GoParser) Parse(contents []byte) ([]string, error) {
	for _, r := range string(contents) {
		g.ingest(r)
	}

	rawStrings := g.strings
	cache := map[string]bool{}
	checkedStrings := []string{}

	for _, str := range rawStrings {
		str = strings.Trim(str, " \t\n\r")
		if str != "" {
			if _, ok := cache[str]; !ok {
				cache[str] = true
				checkedStrings = append(checkedStrings, str)
			}
		}
	}

	return checkedStrings, nil
}

func (g *GoParser) inComment() bool {
	return g.inSingleLineComment || g.inMultiLineComment
}

func (g *GoParser) inString() bool {
	return g.inBacktickedString || g.inDoubleQuotedString
}

func (g *GoParser) updateStrings(r rune) {
	if len(g.strings) == 0 {
		g.strings = []string{string(r)}
	} else {
		g.strings[len(g.strings)-1] = fmt.Sprintf("%s%s", g.strings[len(g.strings)-1], string(r))
	}
}

func (g *GoParser) endString() {
	g.strings = append(g.strings, "")
}

func (g *GoParser) ingest(r rune) {
	switch r {
	case '/':
		if g.inString() {
			g.updateStrings(r)
		} else {
			if g.inMultiLineComment {
				if g.previousRune == '*' {
					g.inMultiLineComment = false
				}
			} else {
				if g.previousRune == '/' {
					g.inSingleLineComment = true
				}
			}
		}
		break
	case '\n':
		if g.inString() {
			g.updateStrings(r)
		} else if g.inSingleLineComment {
			g.inSingleLineComment = false
		}
	case '"':
		if !g.inComment() && !g.inChar {
			if !g.inBacktickedString && g.previousRune != '\\' {
				if g.inDoubleQuotedString {
					g.inDoubleQuotedString = false
					g.endString()
				} else {
					g.inDoubleQuotedString = true
				}
			} else {
				g.updateStrings(r)
			}
		}
	case '`':
		if !g.inComment() && !g.inChar {
			if !g.inDoubleQuotedString {
				if g.inBacktickedString {
					g.inBacktickedString = false
					g.endString()
				} else {
					g.inBacktickedString = true
				}
			}
		}
	case '*':
		if !g.inString() {
			if !g.inSingleLineComment && g.previousRune == '/' {
				g.inMultiLineComment = true
			}
		} else {
			g.updateStrings(r)
		}
	case '\'':
		if !g.inString() && !g.inComment() {
			if g.inChar {
				g.inChar = false
			} else {
				g.inChar = true
			}
		}
	default:
		if g.inString() {
			g.updateStrings(r)
		}
	}

	g.previousRune = r
}
