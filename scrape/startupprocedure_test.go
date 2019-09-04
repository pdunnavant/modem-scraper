package scrape

import (
	"os"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
)

func TestScrapeStartupProcedure(t *testing.T) {
	doc := getConnectionStatusDocumentFromTestFile(t)

	expected := StartupProcedure{
		AcquireDownstreamChannel: Status{
			Status:  "519000000 Hz",
			Comment: "Locked",
		},
		ConnectivityState: Status{
			Status:  "OK",
			Comment: "Operational",
		},
		BootState: Status{
			Status:  "OK",
			Comment: "Operational",
		},
		ConfigurationFile: Status{
			Status:  "OK",
			Comment: "",
		},
		Security: Status{
			Status:  "Enabled",
			Comment: "BPI+",
		},
		DOCSISNetworkAccessEnabled: Status{
			Status:  "Allowed",
			Comment: "",
		},
	}

	actual := scrapeStartupProcedure(doc)
	assert.NotNil(t, actual)
	assert.Equal(t, expected, actual)
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
