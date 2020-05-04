package scrape

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/influxdata/influxdb1-client" // this is important because of a bug in go mod
	client "github.com/influxdata/influxdb1-client/v2"
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

// ToInfluxPoints converts UpstreamBondedChannel to "points"
func (u UpstreamBondedChannel) ToInfluxPoints() ([]*client.Point, error) {
	var points []*client.Point

	channelString := strconv.Itoa(u.Channel)
	channelIDString := strconv.Itoa(u.ChannelID)
	tags := map[string]string{
		"channel":    channelString,
		"channel_id": channelIDString,
	}
	fields := map[string]interface{}{
		"lock_status":     u.LockStatus,
		"us_channel_type": u.USChannelType,
		"frequency_hz":    u.FrequencyHz,
		"width_hz":        u.WidthHz,
		"power_dbmv":      u.PowerdBmV,
	}
	point, err := client.NewPoint("upstream_bonded_channel", tags, fields, time.Now())
	if err != nil {
		return nil, fmt.Errorf("error generating points data for UpstreamBondedChannel: %s", err.Error())
	}

	points = append(points, point)

	return points, nil
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
