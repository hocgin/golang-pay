package wxpay

import (
	"encoding/xml"
	"io"
)

type xmlMapEntry struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}
type StringMap map[string]interface{}

func (m *StringMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*m = StringMap{}
	for {
		var e xmlMapEntry

		err := d.Decode(&e)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		(*m)[e.XMLName.Local] = e.Value
	}
	return nil
}
