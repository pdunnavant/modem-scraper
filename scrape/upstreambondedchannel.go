package scrape

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// UpstreamBondedChannel holds all info from the
// "Upstream Bonded Channels" table on
// /cmconnectionstatus.html.
type UpstreamBondedChannel struct {
	Channel       int
	ChannelID     int
	LockStatus    string
	USChannelType string
	FrequencyHz   int
	WidthHz       int
	PowerdBmV     float64
}

const upstreamBondedChannelTableSelector = "#bg3 > div.container > div.content > center:nth-child(8) > table"

func scrapeUpstreamBondedChannels(doc *goquery.Document) []UpstreamBondedChannel {
	table := doc.Find(upstreamBondedChannelTableSelector)
	tableTbody := table.Children()
	tableTbodyRows := tableTbody.Children()

	upstreamBondedChannels := []UpstreamBondedChannel{}
	tableTbodyRows.Each(func(index int, row *goquery.Selection) {
		// Skip the "title" row as well as the "header" row.
		// These are both regular old <tr> rows on this page.
		if index > 1 {
			upstreamBondedChannels = append(upstreamBondedChannels, makeUpstreamBondedChannel(row))
		}
	})

	return upstreamBondedChannels
}

func makeUpstreamBondedChannel(selection *goquery.Selection) UpstreamBondedChannel {
	rowData := selection.Children()

	upstreamBondedChannel := UpstreamBondedChannel{
		Channel:       getIntRowData(rowData, 0),
		ChannelID:     getIntRowData(rowData, 1),
		LockStatus:    rowData.Get(2).FirstChild.Data,
		USChannelType: rowData.Get(3).FirstChild.Data,
		FrequencyHz:   getIntRowData(rowData, 4),
		WidthHz:       getIntRowData(rowData, 5),
		PowerdBmV:     getFloatRowData(rowData, 6),
	}

	return upstreamBondedChannel
}

func getIntRowData(selection *goquery.Selection, index int) int {
	data := selection.Get(index).FirstChild.Data
	data = strings.Split(data, " ")[0]
	dataInt, _ := strconv.Atoi(data)

	return dataInt
}

func getFloatRowData(selection *goquery.Selection, index int) float64 {
	data := selection.Get(index).FirstChild.Data
	data = strings.Split(data, " ")[0]
	dataFloat, _ := strconv.ParseFloat(data, 64)

	return dataFloat
}
