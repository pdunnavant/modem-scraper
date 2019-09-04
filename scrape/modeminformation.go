package scrape

import (
	"encoding/json"
)

// ModemInformation holds all information from the
// SB8200 status pages.
type ModemInformation struct {
	ConnectionStatus    ConnectionStatus
	SoftwareInformation SoftwareInformation
}

// ToJSON converts ModemInformation to JSON string.
func (m ModemInformation) ToJSON() (string, error) {
	jsonBytes, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	jsonString := string(jsonBytes)
	return jsonString, nil
}

// ToLineProtocol converts ModemInformation to "line
// protocol" for use with InfluxDB.
// TODO: Really, this should be broken up and implemented in:
// - downstreambondedchannels ([]string)
//   - downstreambondedchannel,channel_id=# channel_id=#,frequency=#,etc
// - upstreambondedchannels ([]string)
// - softwareinformation
// - startupprocedure
//
// or maybe... do the above, and also have this (ModemInformation)
// call each of the above to aggregate them all. Then this can
// simply return []string with all the "lines" and they can
// all be sent to InfluxDB at once.
func (m ModemInformation) ToLineProtocol() ([]string, error) {
	// TODO
	return []string{}, nil
}
