package data_types

import (
	"encoding/xml"
)

type DataXML struct {
	XMLName xml.Name `xml:"news"`
	Items   []Item   `xml:"Item"`
}
