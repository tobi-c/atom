package atom

import (
	"encoding/xml"
	"time"
)

// The Atom Syndication Format
// http://www.ietf.org/rfc/rfc4287.txt

const NameSpace = "http://www.w3.org/2005/Atom"

type CommonAttributes struct {
	XMLBase string     `xml:"xml:base,attr,omitempty"`
	XMLLang string     `xml:"xml:lang,attr,omitempty"`
	Attrs   []xml.Attr `xml:",any,attr"`
}

type Feed struct {
	XMLName     xml.Name   `xml:"feed"`
	Author      []Person   `xml:"author,omitempty"`
	Category    []Category `xml:"category,omitempty"`
	Contributor []Person   `xml:"contributor,omitempty"`
	Generator   *Generator `xml:"generator,omitempty"`
	Icon        *URI       `xml:"icon,omitempty"`
	Id          URI        `xml:"id"`
	Link        []Link     `xml:"link,omitempty"`
	Logo        *URI       `xml:"logo,omitempty"`
	Rights      *Text      `xml:"rights,omitempty"`
	Subtitle    *Text      `xml:"subtitle,omitempty"`
	Title       Text       `xml:"title"`
	Updated     Datetime   `xml:"updated"`
	Entry       []Entry    `xml:"entry,omitempty"`
	CommonAttributes
}

type Entry struct {
	XMLName     xml.Name   `xml:"entry"`
	Author      []Person   `xml:"author,omitempty"`
	Category    []Category `xml:"category,omitempty"`
	Content     *Content   `xml:"content,omitempty"`
	Contributor []Person   `xml:"contributor,omitempty"`
	Id          URI        `xml:"id"`
	Link        []Link     `xml:"link,omitempty"`
	Published   *Datetime  `xml:"published,omitempty"`
	Rights      *Text      `xml:"rights,omitempty"`
	Source      *Source    `xml:"source,omitempty"`
	Summary     *Text      `xml:"summary,omitempty"`
	Title       Text       `xml:"title"`
	Updated     Datetime   `xml:"updated"`
	CommonAttributes
}

type Content struct {
	Type string `xml:"type,attr,omitempty"`
	Src  string `xml:"src,attr,omitempty"`
	Body string `xml:",innerxml"`
	CommonAttributes
}

type Text struct {
	Type string `xml:"type,attr,omitempty"`
	Body string `xml:",innerxml"`
	CommonAttributes
}

type Person struct {
	Name  string `xml:"name"`
	URI   string `xml:"uri,omitempty"`
	Email string `xml:",email,omitempty"`
	CommonAttributes
}

type Category struct {
	Term   string `xml:"term,attr"`
	Scheme string `xml:"scheme,attr,omitempty"`
	Label  string `xml:"label,attr,omitempty"`
	CommonAttributes
}

type Generator struct {
	URI     string `xml:"uri,attr,omitempty"`
	Version string `xml:"version,attr,omitempty"`
	Body    string `xml:",chardata"`
	CommonAttributes
}

type URI struct {
	URI string `xml:",chardata"`
	CommonAttributes
}

type UndefinedContent struct {
	Content string `xml:",innerxml"`
}

type Link struct {
	Href     string `xml:"href,attr"`
	Rel      string `xml:"rel,attr,omitempty"`
	Type     string `xml:"type,attr,omitempty"`
	Hreflang string `xml:"hreflang,attr,omitempty"`
	Length   string `xml:"length,attr,omitempty"`
	UndefinedContent
	CommonAttributes
}

type Source struct {
	Author      []Person   `xml:"author,omitempty"`
	Category    []Category `xml:"category,omitempty"`
	Contributor []Person   `xml:"contributor,omitempty"`
	Generator   Generator  `xml:"generator,omitempty"`
	Icon        URI        `xml:"icon,omitempty"`
	Id          URI        `xml:"id,omitempty"`
	Link        []Link     `xml:"link,omitempty"`
	Logo        URI        `xml:"logo,omitempty"`
	Rights      Text       `xml:"rights,omitempty"`
	Subtitle    Text       `xml:"subtitle,omitempty"`
	Title       Text       `xml:"title,omitempty"`
	Updated     Datetime   `xml:"updated,omitempty"`
	CommonAttributes
}

type Datetime struct {
	Datetime time.Time `xml:",chardata"`
	CommonAttributes
}

func (dt *Datetime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	data := struct {
		Datetime string `xml:",chardata"`
		CommonAttributes
	}{
		dt.Datetime.Format(time.RFC3339),
		dt.CommonAttributes,
	}
	return e.EncodeElement(data, start)
}

func (dt *Datetime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	data := struct {
		Datetime string `xml:",chardata"`
		CommonAttributes
	}{}

	err := d.DecodeElement(&data, &start)
	if err != nil {
		return err
	}

	t, err := time.Parse(time.RFC3339, data.Datetime)
	if err != nil {
		return err
	}

	*dt = Datetime{t, data.CommonAttributes}

	return nil
}

func (f *Feed) ToXML() ([]byte, error) {
	buf, err := xml.Marshal(f)
	if err != nil {
		return nil, err
	}
	return append([]byte(xml.Header), buf...), nil
}
