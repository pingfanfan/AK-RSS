package opml

import (
	"encoding/xml"
	"os"
)

type Feed struct {
	Title  string
	XMLURL string
}

type opmlDoc struct {
	Body opmlBody `xml:"body"`
}

type opmlBody struct {
	Outlines []outline `xml:"outline"`
}

type outline struct {
	Text     string    `xml:"text,attr"`
	Title    string    `xml:"title,attr"`
	Type     string    `xml:"type,attr"`
	XMLURL   string    `xml:"xmlUrl,attr"`
	Outlines []outline `xml:"outline"`
}

func LoadFeeds(path string) ([]Feed, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var doc opmlDoc
	if err := xml.Unmarshal(b, &doc); err != nil {
		return nil, err
	}
	feeds := make([]Feed, 0)
	for _, o := range doc.Body.Outlines {
		collect(o, &feeds)
	}
	return feeds, nil
}

func collect(o outline, out *[]Feed) {
	if o.XMLURL != "" {
		title := o.Title
		if title == "" {
			title = o.Text
		}
		*out = append(*out, Feed{Title: title, XMLURL: o.XMLURL})
	}
	for _, child := range o.Outlines {
		collect(child, out)
	}
}
