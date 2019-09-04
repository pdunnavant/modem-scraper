package scrape

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"

	// "github.com/kr/pretty"
	"github.com/pdunnavant/modem-scraper/config"
)

// ModemInformation holds all information from the
// SB8200 status pages.
type ModemInformation struct {
	ConnectionStatus    ConnectionStatus
	SoftwareInformation SoftwareInformation
}

// ToJSON converts ModemInformation to JSON string.
func (m ModemInformation) ToJSON() (string, error) {
	jsonBytes, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	jsonString := string(jsonBytes)
	return jsonString, nil
}

// ToLineProtocol converts ModemInformation to "line
// protocol" for use with InfluxDB.
// TODO: Really, this should be broken up and implemented in:
// - downstreambondedchannels ([]string)
//   - downstreambondedchannel,channel_id=# channel_id=#,frequency=#,etc
// - upstreambondedchannels ([]string)
// - softwareinformation
// - startupprocedure
//
// or maybe... do the above, and also have this (ModemInformation)
// call each of the above to aggregate them all. Then this can
// simply return []string with all the "lines" and they can
// all be sent to InfluxDB at once.
func (m ModemInformation) ToLineProtocol() ([]string, error) {
	// TODO
	return []string{}, nil
}

// Scrape scrapes data from the modem.
func Scrape(config config.Configuration) (*ModemInformation, error) {
	connectionStatus, err := scrapeConnectionStatus(config)
	if err != nil {
		return nil, err
	}
	softwareInformation, err := scrapeSoftwareInformation(config)
	if err != nil {
		return nil, err
	}

	modemInformation := ModemInformation{
		ConnectionStatus:    *connectionStatus,
		SoftwareInformation: *softwareInformation,
	}
	// fmt.Printf("%# v", pretty.Formatter(modemInformation))

	return &modemInformation, nil
}

func getDocumentFromURL(url string) (*goquery.Document, error) {
	fmt.Printf("Grabbing [%s]...\n", url)

	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	elapsed := time.Since(start)
	fmt.Printf("Got [%s]. (Took %s.)\n", url, elapsed)

	return doc, nil
}
