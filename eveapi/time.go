package eveapi

import (
	"time"
)

type EveTime struct{ time.Time }

const eveTimeFormat = "20060102 15:04:05"

func (t *EveTime) Parse(s string) error {
	parsed, err := time.Parse(shortForm, v)
	if err != nil {
		return err
	}
	*t = EveTime{parsed}
	return nil
}

func (t *EveTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	return t.Parse(v)
}
