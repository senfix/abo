package generator

import "github.com/senfix/abo/model"

type Header interface {
	Formattable
	Generate(export model.Export) string
}
