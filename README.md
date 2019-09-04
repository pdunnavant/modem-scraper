What does this do?
==========
This application polls two pages on the Arris SB8200 cable modem:
* http://192.168.100.1/cmconnectionstatus.html
* http://192.168.100.1/cmswinfo.html

The data from those pages is populated into structs, and is
then published out to MQTT as well as InfluxDB.

My intent is to use this data in Home Assistant (MQTT) as well
as InfluxDB/Grafana (for graphing metrics over longer periods
of time).

This is currently a work in progress... and was built for my own
use. That said, if it's useful for someone else, then cool beans.

TODO:
* Add unit tests.
* Build and publish docker container automatically.
