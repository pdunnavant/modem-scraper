package scrape

import "github.com/PuerkitoBio/goquery"

// ConnectionStatus holds all info from /cmconnectionstatus.html.
type ConnectionStatus struct {
	StartupProcedure         StartupProcedure
	DownstreamBondedChannels []DownstreamBondedChannel
	UpstreamBondedChannels   []UpstreamBondedChannel
}

func scrapeConnectionStatus(doc *goquery.Document) (*ConnectionStatus, error) {
	connectionStatus := ConnectionStatus{
		StartupProcedure:         scrapeStartupProcedure(doc),
		DownstreamBondedChannels: scrapeDownstreamBondedChannels(doc),
		UpstreamBondedChannels:   scrapeUpstreamBondedChannels(doc),
	}

	return &connectionStatus, nil
}
