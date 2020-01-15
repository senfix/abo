package generator

import "github.com/senfix/abo/model"

type Line interface {
	Formattable
	Generate(header model.Header, transaction model.Transaction) string
}
