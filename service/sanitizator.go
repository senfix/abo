package service

import (
	"unicode"

	"github.com/senfix/abo/generator"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func sanitize(g generator.Formattable, part string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	result, _, _ := transform.String(t, part)

	length, spacer, position := g.Restrictions()
	if position == generator.SpacerLeft {
		result = padLeft(result, length, spacer)
	}
	if position == generator.SpacerRight {
		result = padRight(result, length, spacer)
	}

	return result
}

func padRight(value string, length int, pad string) string {
	for {
		value += pad
		if len(value) > length {
			return value[0:length]
		}
	}
}

func padLeft(value string, length int, pad string) string {
	for {
		value = pad + value
		if len(value) > length {
			return value[len(value)-length : len(value)]
		}
	}
}
