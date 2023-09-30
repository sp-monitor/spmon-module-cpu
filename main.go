package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func runCommand() ([]byte, error) {
	percentage, err := cpu.Percent(0, true)
	if err != nil {
			return nil, err
	}

	// Create a map to hold the CPU usage data
	data := make(map[string]interface{})
	data["module"] = "cpu"

	// Create a map to hold the CPU usage percentages
	percentages := make(map[string]float64)
	var total float64
	for idx, percent := range percentage {
			// Add the CPU usage percentage for each core to the map
			percentages[fmt.Sprintf("cpu-%d", idx)] = percent
			total += percent
	}
	percentages["avg"] = total / float64(len(percentage))
	// Add the CPU usage percentages to the data map
	data["data"] = []map[string]float64{percentages}

	// Convert the data map to a JSON string
	jsonData, err := json.Marshal(data)
	if err != nil {
			return nil, err
	}
	

	return jsonData, nil
}

func sendJSONData(conn net.Conn, jsonData []byte) error {
	// Send the JSON data over UDP
	_, err := conn.Write(jsonData)
	if err != nil {
		return err
	}
	
	return nil
}

func main() {
	// Define command line flags for the server hostname and port
	serverHost := flag.String("serverHost", "localhost", "the hostname of the remote server")
	serverPort := flag.Int("serverPort", 9559, "the port number of the remote server")
	shorthandHost := flag.String("h", "localhost", "shorthand for serverHost")
	shorthandPort := flag.Int("p", 9559, "shorthand for serverPort")
	serverKey := flag.String("serverKey", "1234567890", "the key to authenticate with the server")
	shorthandKey := flag.String("k", "1234567890", "the key to authenticate with the server")

	flag.Parse()

	// Use the shorthand flags if they are set
	if *serverHost == "localhost"{
			*serverHost = *shorthandHost
	}

	if *serverPort == 9559 {
		*serverPort = *shorthandPort
	}

	if *serverKey == "1234567890" {
		*serverKey = *shorthandKey
	}


	// Create a UDP connection to the remote server
	conn, err := net.Dial("udp", fmt.Sprintf("%s:%d", *serverHost, *serverPort))
	if err != nil {
		fmt.Println("error dial")
		log.Fatalf("Error creating UDP connection: %v", err)
	}
	defer conn.Close()

	for {
			jsonData, err := runCommand()
			if err != nil {
					log.Printf("Error getting CPU usage: %v", err)
					continue
			}

			err = sendJSONData(conn, jsonData)
			if err != nil {
					log.Printf("Error sending JSON data: %v", err)
					continue
			}

			time.Sleep(1 * time.Second)
	}
}