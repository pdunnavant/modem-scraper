package scrape

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"

	// "github.com/kr/pretty"
	"github.com/pdunnavant/modem-scraper/config"
)

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
