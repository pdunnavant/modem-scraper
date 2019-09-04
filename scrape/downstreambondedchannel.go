package scrape

import (
	"github.com/PuerkitoBio/goquery"
)

// DownstreamBondedChannel holds all info from the
// "Downstream Bonded Channels" table on
// /cmconnectionstatus.html.
type DownstreamBondedChannel struct {
	ChannelID      int
	LockStatus     string
	Modulation     string
	FrequencyHz    int
	PowerdBmV      float64
	SNRdB          float64
	Corrected      int
	Uncorrectables int
}

const downstreamBondedChannelTableSelector = "#bg3 > div.container > div.content > center:nth-child(5) > table"

func scrapeDownstreamBondedChannels(doc *goquery.Document) []DownstreamBondedChannel {
	downstreamBondedChannelTable := doc.Find(downstreamBondedChannelTableSelector)
	downstreamBondedChannelTableTbody := downstreamBondedChannelTable.Children()
	downstreamBondedChannelTableTbodyRows := downstreamBondedChannelTableTbody.Children()

	downstreamBondedChannels := []DownstreamBondedChannel{}
	downstreamBondedChannelTableTbodyRows.Each(func(index int, row *goquery.Selection) {
		// Skip the "title" row as well as the "header" row.
		// These are both regular old <tr> rows on this page.
		if index > 1 {
			downstreamBondedChannels = append(downstreamBondedChannels, makeDownstreamBondedChannel(row))
		}
	})

	return downstreamBondedChannels
}

func makeDownstreamBondedChannel(selection *goquery.Selection) DownstreamBondedChannel {
	rowData := selection.Children()
	downstreamBondedChannel := DownstreamBondedChannel{
		ChannelID:      getIntRowData(rowData, 0),
		LockStatus:     rowData.Get(1).FirstChild.Data,
		Modulation:     rowData.Get(2).FirstChild.Data,
		FrequencyHz:    getIntRowData(rowData, 3),
		PowerdBmV:      getFloatRowData(rowData, 4),
		SNRdB:          getFloatRowData(rowData, 5),
		Corrected:      getIntRowData(rowData, 6),
		Uncorrectables: getIntRowData(rowData, 7),
	}

	return downstreamBondedChannel
}
