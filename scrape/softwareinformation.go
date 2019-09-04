package scrape

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/influxdata/influxdb1-client" // this is important because of a bug in go mod
	client "github.com/influxdata/influxdb1-client/v2"
)

// SoftwareInformation holds data pulled from the /cmswinfo.html page.
type SoftwareInformation struct {
	StandardSpecificationCompliant string
	HardwareVersion                string
	SoftwareVersion                string
	MACAddress                     string
	SerialNumber                   string
	UptimeMins                     int
	UptimeString                   string
}

// ToInfluxPoints converts SoftwareInformation to "points"
func (s SoftwareInformation) ToInfluxPoints() ([]*client.Point, error) {
	var points []*client.Point

	// No tags for this specific struct.
	tags := map[string]string{}
	fields := map[string]interface{}{
		"standard_specification_compliant": s.StandardSpecificationCompliant,
		"hardware_version":                 s.HardwareVersion,
		"software_version":                 s.SoftwareVersion,
		"mac_address":                      s.MACAddress,
		"serial_number":                    s.SerialNumber,
		"uptime_mins":                      s.UptimeMins,
		"uptime_string":                    s.UptimeString,
	}
	point, err := client.NewPoint("software_information", tags, fields, time.Now())
	if err != nil {
		return nil, fmt.Errorf("error generating points data for SoftwareInformation: %s", err.Error())
	}

	points = append(points, point)

	return points, nil
}

const standardSpecificationCompliantSelector = "#bg3 > div.container > div.content > table:nth-child(2) > tbody > tr:nth-child(2) > td:nth-child(2)"
const hardwareVersionSelector = "#bg3 > div.container > div.content > table:nth-child(2) > tbody > tr:nth-child(3) > td:nth-child(2)"
const softwareVersionSelector = "#bg3 > div.container > div.content > table:nth-child(2) > tbody > tr:nth-child(4) > td:nth-child(2)"
const macAddressSelector = "#bg3 > div.container > div.content > table:nth-child(2) > tbody > tr:nth-child(5) > td:nth-child(2)"
const serialNumberSelector = "#bg3 > div.container > div.content > table:nth-child(2) > tbody > tr:nth-child(6) > td:nth-child(2)"
const uptimeSelector = "#bg3 > div.container > div.content > table:nth-child(5) > tbody > tr:nth-child(2) > td:nth-child(2)"

func scrapeSoftwareInformation(doc *goquery.Document) *SoftwareInformation {
	uptimeString := doc.Find(uptimeSelector).Text()
	uptimeMins := uptimeToMinutes(uptimeString)
	softwareInformation := SoftwareInformation{
		StandardSpecificationCompliant: doc.Find(standardSpecificationCompliantSelector).Text(),
		HardwareVersion:                doc.Find(hardwareVersionSelector).Text(),
		SoftwareVersion:                doc.Find(softwareVersionSelector).Text(),
		MACAddress:                     doc.Find(macAddressSelector).Text(),
		SerialNumber:                   doc.Find(serialNumberSelector).Text(),
		UptimeMins:                     uptimeMins,
		UptimeString:                   uptimeString,
	}

	return &softwareInformation
}

// "0 days 02h:44m:31s.00"
func uptimeToMinutes(uptime string) int {
	daysString := strings.Split(uptime, " ")[0]
	timeString := strings.Split(uptime, " ")[2]
	timeString = strings.ReplaceAll(timeString, "h", "")
	timeString = strings.ReplaceAll(timeString, "m", "")
	hoursString := strings.Split(timeString, ":")[0]
	minutesString := strings.Split(timeString, ":")[1]

	days, _ := strconv.Atoi(daysString)
	hours, _ := strconv.Atoi(hoursString)
	minutes, _ := strconv.Atoi(minutesString)

	totalMinutes := (days * 24 * 60) + (hours * 60) + minutes
	return totalMinutes
}
