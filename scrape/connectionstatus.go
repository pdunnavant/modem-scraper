package scrape

import (
	"github.com/PuerkitoBio/goquery"
	_ "github.com/influxdata/influxdb1-client" // this is important because of a bug in go mod
	client "github.com/influxdata/influxdb1-client/v2"
)

// ConnectionStatus holds all info from /cmconnectionstatus.html.
type ConnectionStatus struct {
	StartupProcedure         StartupProcedure
	DownstreamBondedChannels []DownstreamBondedChannel
	UpstreamBondedChannels   []UpstreamBondedChannel
}

// ToInfluxPoints converts ConnectionStatus to "points"
func (c ConnectionStatus) ToInfluxPoints() ([]*client.Point, error) {
	var points []*client.Point

	influxPoints, err := c.StartupProcedure.ToInfluxPoints()
	if err != nil {
		return nil, err
	}
	points = append(points, influxPoints...)

	influxPoints, err = buildDownstreamBondedChannelPoints(c.DownstreamBondedChannels)
	if err != nil {
		return nil, err
	}
	points = append(points, influxPoints...)

	influxPoints, err = buildUpstreamBondedChannelPoints(c.UpstreamBondedChannels)
	if err != nil {
		return nil, err
	}
	points = append(points, influxPoints...)

	return points, nil
}

func scrapeConnectionStatus(doc *goquery.Document) *ConnectionStatus {
	connectionStatus := ConnectionStatus{
		StartupProcedure:         scrapeStartupProcedure(doc),
		DownstreamBondedChannels: scrapeDownstreamBondedChannels(doc),
		UpstreamBondedChannels:   scrapeUpstreamBondedChannels(doc),
	}

	return &connectionStatus
}

func buildDownstreamBondedChannelPoints(channels []DownstreamBondedChannel) ([]*client.Point, error) {
	var points []*client.Point

	for _, channel := range channels {
		influxPoints, err := channel.ToInfluxPoints()
		if err != nil {
			return nil, err
		}
		points = append(points, influxPoints...)
	}

	return points, nil
}

func buildUpstreamBondedChannelPoints(channels []UpstreamBondedChannel) ([]*client.Point, error) {
	var points []*client.Point

	for _, channel := range channels {
		influxPoints, err := channel.ToInfluxPoints()
		if err != nil {
			return nil, err
		}
		points = append(points, influxPoints...)
	}

	return points, nil
}
