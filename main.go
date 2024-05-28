package main

import (
    "flag"
    "fmt"
    "os"

    MQTT "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
    fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler MQTT.OnConnectHandler = func(client MQTT.Client) {
    fmt.Println("Connected")
}

var connectLostHandler MQTT.ConnectionLostHandler = func(client MQTT.Client, err error) {
    fmt.Printf("Connect lost: %v", err)
}

func main() {
    server := flag.String("server", "", "The MQTT server address (e.g., tcp://localhost:1883)")
    username := flag.String("username", "", "The username for the MQTT server")
    password := flag.String("password", "", "The password for the MQTT server")
    topic := flag.String("topic", "test/topic", "The topic to publish/subscribe to")
    message := flag.String("message", "Hello, MQTT from Go with Auth!", "The message to publish")

    flag.Parse()

    if *server == "" || *username == "" || *password == "" || *topic == "" || *message == "" {
        fmt.Println("All arguments (server, username, password, topic, message) are required.")
        flag.Usage()
        os.Exit(1)
    }

    opts := MQTT.NewClientOptions()
    opts.AddBroker(*server)
    opts.SetClientID("go_mqtt_client")
    opts.SetUsername(*username)
    opts.SetPassword(*password)
    opts.SetDefaultPublishHandler(messagePubHandler)
    opts.OnConnect = connectHandler
    opts.OnConnectionLost = connectLostHandler

    client := MQTT.NewClient(opts)
    if token := client.Connect(); token.Wait() && token.Error() != nil {
        fmt.Println(token.Error())
        os.Exit(1)
    }

    // Publish a message
    token := client.Publish(*topic, 0, false, *message)
    token.Wait()

    client.Disconnect(250)
}
