package scrape

import "github.com/pdunnavant/modem-scraper/config"

// ConnectionStatus holds all info from /cmconnectionstatus.html.
type ConnectionStatus struct {
	StartupProcedure         StartupProcedure
	DownstreamBondedChannels []DownstreamBondedChannel
	UpstreamBondedChannels   []UpstreamBondedChannel
}

func scrapeConnectionStatus(config config.Configuration) (*ConnectionStatus, error) {
	doc, err := getDocumentFromURL(config.IP + "/cmconnectionstatus.html")
	if err != nil {
		return nil, err
	}

	connectionStatus := ConnectionStatus{
		StartupProcedure:         scrapeStartupProcedure(doc),
		DownstreamBondedChannels: scrapeDownstreamBondedChannels(doc),
		UpstreamBondedChannels:   scrapeUpstreamBondedChannels(doc),
	}

	return &connectionStatus, nil
}
