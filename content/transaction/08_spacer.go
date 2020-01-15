package transaction

import (
	"github.com/senfix/abo/generator"
	"github.com/senfix/abo/model"
)

//Typ záznamu – detail transakce, uveden text 075
type Spacer struct {
}

func NewSpacer() generator.Line {
	return &Spacer{}
}

func (t *Spacer) Generate(header model.Header, transfer model.Transaction) string {
	return "00"
}

func (t *Spacer) Restrictions() (length int, spacer string, position generator.SpacerPosition) {
	return 2, "0", generator.SpacerLeft
}
