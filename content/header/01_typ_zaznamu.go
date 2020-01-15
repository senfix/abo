package header

import (
	"github.com/senfix/abo/constants"
	"github.com/senfix/abo/generator"
	"github.com/senfix/abo/model"
)

//Typ záznamu – hlavička výpisu, uveden text 074
type TypZaznamu struct {
}

func NewTypZaznamu() generator.Header {
	return &TypZaznamu{}
}

func (t *TypZaznamu) Generate(export model.Export) string {
	return constants.CODE_HEADER
}

func (t *TypZaznamu) Restrictions() (length int, spacer string, position generator.SpacerPosition) {
	return 3, "0", generator.SpacerLeft
}
