package transaction

import (
	"github.com/senfix/abo/constants"
	"github.com/senfix/abo/generator"
	"github.com/senfix/abo/model"
)

//Typ záznamu – detail transakce, uveden text 075
type TypZaznamu struct {
}

func NewTypZaznamu() generator.Line {
	return &TypZaznamu{}
}

func (t *TypZaznamu) Generate(header model.Header, transfer model.Transaction) string {
	return constants.CODE_TRANSACTION
}

func (t *TypZaznamu) Restrictions() (length int, spacer string, position generator.SpacerPosition) {
	return 3, "0", generator.SpacerLeft
}
