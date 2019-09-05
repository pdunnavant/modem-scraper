package influxdb

import (
	"fmt"
	"time"

	_ "github.com/influxdata/influxdb1-client" // this is important because of a bug in go mod
	client "github.com/influxdata/influxdb1-client/v2"
	"github.com/pdunnavant/modem-scraper/config"
	"github.com/pdunnavant/modem-scraper/scrape"
)

// Publish publishes the data within modemInformation to
// the InfluxDB server configuration within the given
// configuration.
func Publish(config config.InfluxDB, modemInformation scrape.ModemInformation) error {
	start := time.Now()

	addr := makeAddr(config.Hostname, config.Port)

	fmt.Printf("Connecting to InfluxDB server [%s]...\n", addr)
	influx, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     addr,
		Username: config.Username,
		Password: config.Password,
	})
	if err != nil {
		return fmt.Errorf("error creating InfluxDB client: %s", err.Error())
	}
	defer influx.Close()

	batchPoints, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  config.Database,
		Precision: "s",
	})
	points, err := modemInformation.ToInfluxPoints()
	if err != nil {
		return err
	}
	batchPoints.AddPoints(points)

	fmt.Printf("Writing [%d] data points to InfluxDB database [%s]...\n", len(points), config.Database)
	err = influx.Write(batchPoints)
	if err != nil {
		return fmt.Errorf("error writing data to InfluxDB: %s", err.Error())
	}

	elapsed := time.Since(start)
	fmt.Printf("Finished writing to InfluxDB. (Took %s.)\n", elapsed)

	return nil
}

func makeAddr(hostname string, port string) string {
	// TODO: allow specifying useSsl in config
	return fmt.Sprintf("http://%s:%s", hostname, port)
}

func buildPoints() []*client.Point {
	var points []*client.Point

	// TODO:
	// - build a point for each dbc
	// - build a point for each ubc
	// - build a point for startupprocedure data
	// - build a point for softwareinformation data
	//
	// Do the above with a single call to
	// ModemInformation.ToInfluxPoints(), which
	// will itself do all the work of calling the
	// downstream things, etc, etc.
	//
	// These should all use the same time.Now().

	tags := map[string]string{"cpu": "cpu-total"}
	fields := map[string]interface{}{
		"idle":   10.1,
		"system": 53.3,
		"user":   46.6,
	}
	point, err := client.NewPoint("cpu_usage", tags, fields, time.Now())
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	points = append(points, point)

	return points
}
