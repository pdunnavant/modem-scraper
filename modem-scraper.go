package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/pdunnavant/modem-scraper/config"
	"github.com/pdunnavant/modem-scraper/influxdb"
	"github.com/pdunnavant/modem-scraper/mqtt"
	"github.com/pdunnavant/modem-scraper/scrape"
	"github.com/robfig/cron"
	"github.com/spf13/viper"
)

// BuildVersion is the version of the binary, and is set with ldflags at build time.
var BuildVersion = "UNKNOWN"

// CliInputs holds the data passed in via CLI parameters
type CliInputs struct {
	BuildVersion string
	Config       string
	ShowVersion  bool
}

func main() {
	cliInputs := CliInputs{
		BuildVersion: BuildVersion,
	}
	flags := flag.NewFlagSet("modem-scraper", 0)
	flags.StringVar(&cliInputs.Config, "config", "config.yaml", "Set the location for the YAML config file")
	flags.BoolVar(&cliInputs.ShowVersion, "version", false, "Print the version of modem-script")
	flags.Parse(os.Args[1:])

	if cliInputs.ShowVersion {
		fmt.Println(cliInputs.BuildVersion)
		os.Exit(0)
	}

	configuration, err := parseConfiguration(cliInputs.Config)
	if err != nil {
		log.Fatalf(err.Error())
	}

	c := cron.New()
	c.AddFunc(configuration.PollSchedule, func() {
		fmt.Println("Waking up...")
		modemInformation, err := scrape.Scrape(*configuration)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		err = mqtt.Publish(configuration.MQTT, *modemInformation)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		err = influxdb.Publish(configuration.InfluxDB, *modemInformation)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("Going back to sleep...")
	})
	go c.Start()

	// Wait forever, but just for an OS interrupt/kill.
	fmt.Println("Started.")
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}

func parseConfiguration(configPath string) (*config.Configuration, error) {
	viper.SetConfigFile(configPath)
	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("Error reading config file, %s", err)
	}

	var configuration config.Configuration
	err := viper.Unmarshal(&configuration)
	if err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %s", err)
	}

	// fmt.Printf("%# v", pretty.Formatter(configuration))
	return &configuration, nil
}
