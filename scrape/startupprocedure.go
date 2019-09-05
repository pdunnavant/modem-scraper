package scrape

import (
	"fmt"
	"time"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/influxdata/influxdb1-client" // this is important because of a bug in go mod
	client "github.com/influxdata/influxdb1-client/v2"
)

// StartupProcedure holds the data from the
// "Startup Procedure" table on
// /cmconnectionstatus.html.
type StartupProcedure struct {
	AcquireDownstreamChannel   Status
	ConnectivityState          Status
	BootState                  Status
	ConfigurationFile          Status
	Security                   Status
	DOCSISNetworkAccessEnabled Status
}

// Status holds startup procedure status/comment.
type Status struct {
	Status  string
	Comment string
}

// ToInfluxPoints converts StartupProcedure to "points"
func (s StartupProcedure) ToInfluxPoints() ([]*client.Point, error) {
	var points []*client.Point

	// No tags for this specific struct.
	tags := map[string]string{}
	fields := map[string]interface{}{
		"acquire_downstream_channel_status":     s.AcquireDownstreamChannel.Status,
		"acquire_downstream_channel_comment":    s.AcquireDownstreamChannel.Comment,
		"connectivity_state_status":             s.ConnectivityState.Status,
		"connectivity_state_comment":            s.ConnectivityState.Comment,
		"boot_state_status":                     s.BootState.Status,
		"boot_state_comment":                    s.BootState.Comment,
		"configuration_file_status":             s.ConfigurationFile.Status,
		"configuration_file_comment":            s.ConfigurationFile.Comment,
		"security_status":                       s.Security.Status,
		"security_comment":                      s.Security.Comment,
		"docsis_network_access_enabled_status":  s.DOCSISNetworkAccessEnabled.Status,
		"docsis_network_access_enabled_comment": s.DOCSISNetworkAccessEnabled.Comment,
	}
	point, err := client.NewPoint("startup_procedure", tags, fields, time.Now())
	if err != nil {
		return nil, fmt.Errorf("error generating points data for StartupProcedure: %s", err.Error())
	}

	points = append(points, point)

	return points, nil
}

const acquireDownstreamChannelSelector = "#bg3 > div.container > div.content > center:nth-child(2) > table > tbody > tr:nth-child(3) > td:nth-child(2)"
const connectivityStateSelector = "#bg3 > div.container > div.content > center:nth-child(2) > table > tbody > tr:nth-child(4) > td:nth-child(2)"
const bootStateSelector = "#bg3 > div.container > div.content > center:nth-child(2) > table > tbody > tr:nth-child(5) > td:nth-child(2)"
const configurationFileSelector = "#bg3 > div.container > div.content > center:nth-child(2) > table > tbody > tr:nth-child(6) > td:nth-child(2)"
const securitySelector = "#bg3 > div.container > div.content > center:nth-child(2) > table > tbody > tr:nth-child(7) > td:nth-child(2)"
const docsisNetworkAccessEnabledSelector = "#bg3 > div.container > div.content > center:nth-child(2) > table > tbody > tr:nth-child(8) > td:nth-child(2)"

func scrapeStartupProcedure(doc *goquery.Document) StartupProcedure {
	startupProcedure := StartupProcedure{
		AcquireDownstreamChannel:   makeStatus(doc, acquireDownstreamChannelSelector),
		ConnectivityState:          makeStatus(doc, connectivityStateSelector),
		BootState:                  makeStatus(doc, bootStateSelector),
		ConfigurationFile:          makeStatus(doc, configurationFileSelector),
		Security:                   makeStatus(doc, securitySelector),
		DOCSISNetworkAccessEnabled: makeStatus(doc, docsisNetworkAccessEnabledSelector),
	}

	return startupProcedure
}

func makeStatus(doc *goquery.Document, selector string) Status {
	selection := doc.Find(selector)

	status := Status{
		Status:  selection.Text(),
		Comment: selection.Next().Text(),
	}
	return status
}
