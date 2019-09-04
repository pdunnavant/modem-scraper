package influxdb

import (
	"fmt"

	_ "github.com/influxdata/influxdb1-client" // this is important because of a bug in go mod
	client "github.com/influxdata/influxdb1-client/v2"
	"github.com/pdunnavant/modem-scraper/config"
	"github.com/pdunnavant/modem-scraper/scrape"
)

// Publish publishes the data within modemInformation to
// the InfluxDB server configuration within the given
// configuration.
func Publish(config config.InfluxDB, modemInformation scrape.ModemInformation) error {
	influx, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: makeAddr(config.Hostname, config.Port),
	})
	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
	}
	defer influx.Close()

	return nil
}

func makeAddr(hostname string, port string) string {
	return fmt.Sprintf("http://%s:%s", hostname, port)
}
