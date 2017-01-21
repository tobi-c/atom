package atom_test

import (
	atom "."
	"encoding/xml"
	"io/ioutil"
	"testing"
	"time"
)

func TestMarshalFeed(t *testing.T) {
	data := []struct {
		input    atom.Feed
		expected string
	}{
		{
			atom.Feed{},
			`<feed><id></id><title></title><updated>0001-01-01T00:00:00Z</updated></feed>`,
		},
		{
			atom.Feed{
				Author:      []atom.Person{{Name: "Author"}},
				Category:    []atom.Category{{Term: "term"}},
				Contributor: []atom.Person{{Name: "Contributor"}},
				Generator:   &atom.Generator{Body: "Generator"},
				Icon:        &atom.URI{URI: "Icon"},
				Id:          atom.URI{URI: "Id"},
				Link:        []atom.Link{{Href: "example.com"}},
				Logo:        &atom.URI{URI: "Logo"},
				Rights:      &atom.Text{Body: "Rights"},
				Subtitle:    &atom.Text{Body: "Subtitle"},
				Title:       atom.Text{Body: "Title"},
				Updated:     atom.Datetime{Datetime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)},
			},
			`<feed><author><name>Author</name></author><category term="term"></category><contributor><name>Contributor</name></contributor><generator>Generator</generator><icon>Icon</icon><id>Id</id><link href="example.com"></link><logo>Logo</logo><rights>Rights</rights><subtitle>Subtitle</subtitle><title>Title</title><updated>2009-11-10T23:00:00Z</updated></feed>`,
		},
	}

	for _, d := range data {
		buf, err := xml.Marshal(&d.input)
		if err != nil {
			t.Error(err)
		}

		actual := string(buf)
		if d.expected != actual {
			t.Errorf("expected: %v, actual: %v", d.expected, actual)
		}
	}
}

func TestToXML(t *testing.T) {
	data := []struct {
		input    atom.Feed
		expected string
	}{
		{
			atom.Feed{},
			xml.Header + `<feed><id></id><title></title><updated>0001-01-01T00:00:00Z</updated></feed>`,
		},
		{
			atom.Feed{
				Author:      []atom.Person{{Name: "Author"}},
				Category:    []atom.Category{{Term: "term"}},
				Contributor: []atom.Person{{Name: "Contributor"}},
				Generator:   &atom.Generator{Body: "Generator"},
				Icon:        &atom.URI{URI: "Icon"},
				Id:          atom.URI{URI: "Id"},
				Link:        []atom.Link{{Href: "example.com"}},
				Logo:        &atom.URI{URI: "Logo"},
				Rights:      &atom.Text{Body: "Rights"},
				Subtitle:    &atom.Text{Body: "Subtitle"},
				Title:       atom.Text{Body: "Title"},
				Updated:     atom.Datetime{Datetime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)},
			},
			xml.Header + `<feed><author><name>Author</name></author><category term="term"></category><contributor><name>Contributor</name></contributor><generator>Generator</generator><icon>Icon</icon><id>Id</id><link href="example.com"></link><logo>Logo</logo><rights>Rights</rights><subtitle>Subtitle</subtitle><title>Title</title><updated>2009-11-10T23:00:00Z</updated></feed>`,
		},
	}

	for _, d := range data {
		buf, err := d.input.ToXML()
		if err != nil {
			t.Error(err)
		}

		actual := string(buf)
		if d.expected != actual {
			t.Errorf("expected: %v, actual: %v", d.expected, actual)
		}
	}
}

func TestUnmarshalFeedFile(t *testing.T) {
	data := []struct {
		input    string
		expected atom.Feed
	}{
		{
			"./testdata/test1.xml",
			atom.Feed{
				Author:      []atom.Person{{Name: "John Doe"}},
				Category:    []atom.Category{},
				Contributor: []atom.Person{},
				Generator:   &atom.Generator{},
				Icon:        &atom.URI{},
				Id:          atom.URI{URI: "urn:uuid:60a76c80-d399-11d9-b93C-0003939e0af6"},
				Link:        []atom.Link{{Href: "http://example.org/"}},
				Title:       atom.Text{Body: "Example Feed"},
				Updated:     atom.Datetime{Datetime: time.Date(2003, 12, 13, 18, 30, 2, 0, time.UTC)},
			},
		},
		{
			"./testdata/test2.xml",
			atom.Feed{
				Author:      []atom.Person{},
				Category:    []atom.Category{},
				Contributor: []atom.Person{},
				Generator:   &atom.Generator{},
				Icon:        &atom.URI{URI: "Icon"},
				Id:          atom.URI{URI: "tag:example.org,2003:3"},
				Link:        []atom.Link{{Href: "http://example.org/"}},
				Title:       atom.Text{Body: "dive into mark"},
				Updated:     atom.Datetime{Datetime: time.Date(2005, 07, 31, 12, 29, 29, 0, time.UTC)},
			},
		},
	}

	for _, d := range data {
		bytes, err := ioutil.ReadFile(d.input)
		if err != nil {
			t.Error(err)
		}

		actual := atom.Feed{}
		err = xml.Unmarshal(bytes, &actual)
		if err != nil {
			t.Error(err)
		}

		if d.expected.Link[0].Href != actual.Link[0].Href {
			t.Errorf("expected: %v, actual: %v", d.expected, actual)
		}
		if d.expected.Title.Body != actual.Title.Body {
			t.Errorf("expected: %v, actual: %v", d.expected, actual)
		}
		if d.expected.Id.URI != actual.Id.URI {
			t.Errorf("expected: %v, actual: %v", d.expected, actual)
		}
		if d.expected.Updated.Datetime != actual.Updated.Datetime {
			t.Errorf("expected: %v, actual: %v", d.expected, actual)
		}
	}
}
