package mqtt

import (
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/pdunnavant/modem-scraper/config"
	"github.com/pdunnavant/modem-scraper/scrape"
)

// Publish publishes the jsonified modemInformation to
// the MQTT server configuration within the given
// configuration.
func Publish(config config.MQTT, modemInformation scrape.ModemInformation) error {
	broker := makeBroker(config.Hostname, config.Port)

	opts := MQTT.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(config.ClientID)
	opts.SetUsername(config.Username)
	opts.SetPassword(config.Password)

	client := MQTT.NewClient(opts)

	fmt.Printf("Connecting to MQTT server [%s]...\n", broker)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	fmt.Printf("Publishing to topic [%s]...\n", config.Topic)

	payload, err := modemInformation.ToJSON()
	if err != nil {
		return err
	}
	// fmt.Println(payload)
	token := client.Publish(config.Topic, byte(0), false, payload)
	token.Wait()

	client.Disconnect(250)

	fmt.Println("Finished publishing.")

	return nil
}

func makeBroker(hostname string, port string) string {
	return fmt.Sprintf("tcp://%s:%s", hostname, port)
}
