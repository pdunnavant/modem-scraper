package config

// Configuration holds all configuration for modem-scraper.
type Configuration struct {
	ModemType    string
	IP           string
	PollSchedule string
	MQTT         MQTT
	InfluxDB     InfluxDB
}

// MQTT holds MQTT connection configuration.
type MQTT struct {
	Hostname string
	Port     string
	Username string
	Password string
	Topic    string
	ClientID string
}

// InfluxDB holds InfluxDB connection configuration.
type InfluxDB struct {
	Hostname string
	Port     string
	Database string
	Username string
	Password string
}
