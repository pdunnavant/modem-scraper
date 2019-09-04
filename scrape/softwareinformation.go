package scrape

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
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

const standardSpecificationCompliantSelector = "#bg3 > div.container > div.content > table:nth-child(2) > tbody > tr:nth-child(2) > td:nth-child(2)"
const hardwareVersionSelector = "#bg3 > div.container > div.content > table:nth-child(2) > tbody > tr:nth-child(3) > td:nth-child(2)"
const softwareVersionSelector = "#bg3 > div.container > div.content > table:nth-child(2) > tbody > tr:nth-child(4) > td:nth-child(2)"
const macAddressSelector = "#bg3 > div.container > div.content > table:nth-child(2) > tbody > tr:nth-child(5) > td:nth-child(2)"
const serialNumberSelector = "#bg3 > div.container > div.content > table:nth-child(2) > tbody > tr:nth-child(6) > td:nth-child(2)"
const uptimeSelector = "#bg3 > div.container > div.content > table:nth-child(5) > tbody > tr:nth-child(2) > td:nth-child(2)"

func scrapeSoftwareInformation(doc *goquery.Document) (*SoftwareInformation, error) {
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

	return &softwareInformation, nil
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
