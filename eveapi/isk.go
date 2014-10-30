package eveapi

import (
	"encoding/xml"
	"strconv"
)

type ISK struct{ int64 }

func (i *ISK) Parse(s string) error {
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	*i = ISK{v}
	return nil
}

func (t *ISK) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	return t.Parse(v)
}
