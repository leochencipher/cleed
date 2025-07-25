package cleed

import (
	"testing"
	"time"

	"github.com/radulucut/cleed/internal/storage"
)

var (
	defaultCurrentTime = time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
)

func localStorageCleanup(t *testing.T, storage *storage.LocalStorage) {
	err := storage.ClearAll()
	if err != nil {
		t.Fatal(err)
	}
}

type FeedItem struct {
	Title      string
	Link       string
	Published  string
	Categories []string
}

func createRSS(items []*FeedItem) string {
	var itemsStr = ""
	for _, item := range items {
		itemsStr += "<item><title>" +
			item.Title + "</title><link>" +
			item.Link + "</link><pubDate>" +
			item.Published + "</pubDate>"
		for i := range item.Categories {
			itemsStr += "<category>" + item.Categories[i] + "</category>"
		}
		itemsStr += "</item>"
	}
	return `<rss version="2.0"><channel>
<title>RSS Feed</title>
<description>Feed description</description>
<link>https://rss-feed.com/</link>` + itemsStr + "</channel></rss>"
}

func createDefaultRSS() string {
	return `<rss version="2.0">
	<channel>
		<title>RSS Feed</title>
		<description>RSS Feed description</description>
		<link>https://rss-feed.com/</link>
		<item>
			<title>Item 1</title>
			<link>https://rss-feed.com/item-1/</link>
			<pubDate>Wed, 31 Dec 2023 23:45:00 GMT</pubDate>
		</item>
		<item>
			<title>Item 2</title>
			<link>/item-2/</link>
			<pubDate>Sat, 18 May 2019 21:00:00 GMT</pubDate>
		</item>
	</channel>
</rss>`
}

func createDefaultAtom() string {
	return `<?xml version="1.0" encoding="utf-8"?>
<feed xmlns="http://www.w3.org/2005/Atom">
	<title>Atom Feed</title>
	<subtitle>Atom Feed description</subtitle>
	<link href="https://atom-feed.com/"/>
	<entry>
		<title>Item 1</title>
		<link href="https://atom-feed.com/item-1/"/>
		<updated>2023-12-31T06:00:00Z</updated>
	</entry>
	<entry>
		<title>Item 2</title>
		<link href="https://atom-feed.com/item-2/"/>
		<updated>2019-08-20T21:00:00Z</updated>
	</entry>
</feed>`
}
