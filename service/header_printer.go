package service

import (
	"io"

	"github.com/senfix/abo/constants"

	tx "github.com/senfix/abo/content/header"
	"github.com/senfix/abo/generator"
	"github.com/senfix/abo/model"
)

type HeaderPrinter struct {
}

func (t *HeaderPrinter) Export(export model.Export, writer io.Writer) {

	generators := []generator.Header{
		tx.NewTypZaznamu(),
		tx.NewCisloUctu(),
		tx.NewMajitelUctu(),
		tx.NewDatumPocatecnihoZustatku(),
		tx.NewPocatecniZustatek(),
		tx.NewZnamenkoPocatecnihoZustatku(),
		tx.NewKoncovyZustatek(),
		tx.NewZnamenkoKoncovehoZustatku(),
		tx.NewSumaKreditnichPolozek(),
		tx.NewZnamenkoProKreditniPlatby(),
		tx.NewSumaDebetnichPolozek(),
		tx.NewZnamenkoProDebetniPlatby(),
		tx.NewPoradoveCislo(),
		tx.NewDatumVypisu(),
		tx.NewFiller(),
	}

	for _, g := range generators {
		part := g.Generate(export)
		writer.Write([]byte(sanitize(g, part)))
	}
	writer.Write([]byte(constants.LINE_SEPARATOR))
}
