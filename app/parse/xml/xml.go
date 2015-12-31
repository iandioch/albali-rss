package xml

import (
	"fmt"
	"encoding/xml"
)

type Item struct {
	Title string `xml:"title"`
	Link string `xml:"link"`
	Description string `xml:"description"`
}

type Channel struct {
	RSSVersion string `xml:"version,attr"`
	ChannelTitle string `xml:"channel>title"`
	ChannelLink string `xml:"channel>link"`
	ChannelDescription string `xml:"channel>description"`
	Items []Item `xml:"channel>item"`
}

func Parse(raw string) (Channel, error) {
	v := Channel{}
	err := xml.Unmarshal([]byte(raw), &v)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return Channel{}, err
	}
	return v, nil
}