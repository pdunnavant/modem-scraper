package scrape

import (
	"testing"

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
