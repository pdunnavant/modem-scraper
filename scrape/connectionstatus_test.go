package scrape

import (
	"os"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
)

func TestScrapeConnectionStatus(t *testing.T) {
	doc := getConnectionStatusDocumentFromTestFile(t)

	actual := scrapeConnectionStatus(doc)
	assert.NotNil(t, actual)
	assert.NotNil(t, actual.StartupProcedure)
	assert.NotNil(t, actual.DownstreamBondedChannels)
	assert.Len(t, actual.DownstreamBondedChannels, 32)
	assert.NotNil(t, actual.UpstreamBondedChannels)
	assert.Len(t, actual.UpstreamBondedChannels, 4)
}

func getConnectionStatusDocumentFromTestFile(t *testing.T) *goquery.Document {
	filePath := "../testdata/sb8200/cmconnectionstatus.html"
	fileReader, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("unable to open file for reading: [%s]", filePath)
	}

	doc, err := goquery.NewDocumentFromReader(fileReader)
	if err != nil {
		t.Fatalf("unable to generate goquery document from file: [%s]", filePath)
	}

	return doc
}
