package scrape

import (
	"testing"

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
