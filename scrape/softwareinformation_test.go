package scrape

import (
	"os"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
)

func TestUptimeToMinutesWith0d2h44mReturns164(t *testing.T) {
	uptime := "0 days 02h:44m:31s.00"
	expected := 164
	actual := uptimeToMinutes(uptime)
	assert.Equal(t, expected, actual)
}

func TestUptimeToMinutesWith1d0h0mReturns1440(t *testing.T) {
	uptime := "1 days 00h:00m:31s.00"
	expected := 1440
	actual := uptimeToMinutes(uptime)
	assert.Equal(t, expected, actual)
}

func TestUptimeToMinutesWith1d23h59mReturns2879(t *testing.T) {
	uptime := "1 days 23h:59m:31s.00"
	expected := 2879
	actual := uptimeToMinutes(uptime)
	assert.Equal(t, expected, actual)
}

func TestScrapeSoftwareInformation(t *testing.T) {
	doc := getSoftwareInformationDocumentFromTestFile(t)

	expected := &SoftwareInformation{
		StandardSpecificationCompliant: "Docsis 3.1",
		HardwareVersion:                "4",
		SoftwareVersion:                "SB8200.0200.174F.311915.NSH.RT.NA",
		MACAddress:                     "TH:IS:IS:FA:KE:00",
		SerialNumber:                   "THISISFAKE12345",
		UptimeMins:                     2292,
		UptimeString:                   "1 days 14h:12m:38s.00",
	}

	actual := scrapeSoftwareInformation(doc)
	assert.NotNil(t, actual)
	assert.Equal(t, expected, actual)
}

func getSoftwareInformationDocumentFromTestFile(t *testing.T) *goquery.Document {
	filePath := "../testdata/sb8200/cmswinfo.html"
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
