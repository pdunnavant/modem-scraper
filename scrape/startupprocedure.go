package scrape

import "github.com/PuerkitoBio/goquery"

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
