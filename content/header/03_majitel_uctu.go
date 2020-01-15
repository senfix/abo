package header

import (
	"github.com/senfix/abo/generator"
	"github.com/senfix/abo/model"
)

//Majitel účtu (prvních 20 znaků z názvu majitele účtu
type MajitelUctu struct {
}

func NewMajitelUctu() generator.Header {
	return &MajitelUctu{}
}

func (t *MajitelUctu) Generate(export model.Export) string {
	return export.Header.NazevUctu
}

func (t *MajitelUctu) Restrictions() (length int, spacer string, position generator.SpacerPosition) {
	return 20, " ", generator.SpacerRight
}
