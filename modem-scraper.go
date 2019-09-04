package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/pdunnavant/modem-scraper/config"
	"github.com/pdunnavant/modem-scraper/mqtt"
	"github.com/pdunnavant/modem-scraper/scrape"
	"github.com/robfig/cron"
	"github.com/spf13/viper"
)

func main() {
	configuration, err := parseConfiguration()
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

		// fmt.Printf("%# v\n", pretty.Formatter(modemInformation))

		err = mqtt.Publish(configuration.MQTT, *modemInformation)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// TODO:
		// - send to mqtt
		// - send to influxdb

		fmt.Println("Going back to sleep...")
	})
	go c.Start()

	// Wait forever, but just for an OS interrupt/kill.
	fmt.Println("Started.")
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}

func parseConfiguration() (*config.Configuration, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
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
