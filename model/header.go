package model

import "time"

type Header struct {
	CisloUctu         string
	NazevUctu         string
	PocatecniZustatek float64
	ExportOd          time.Time
	ExportDo          time.Time
}
