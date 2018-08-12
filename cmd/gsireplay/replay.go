package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/bergmannf/csgo-gsi/csgogsi"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	GSI_FILE        = "./gsi.json"
	FILE_NOT_READ   = 1
	NO_GSI_RECEIVER = 2
)

type options struct {
	File   *string
	Server *string
}

func parseArgs() options {
	var opts options
	opts.File = flag.String("file", GSI_FILE, "File to replay events from.")
	opts.Server = flag.String("server", "", "Server that will receive events.")
	flag.Parse()
	return opts
}

func repeat(events []csgogsi.GSIData, receiver string) {
	log.Printf("Repeating %d events to: %s", len(events), receiver)
	for n, currentEvent := range events {
		nextEvent := events[n+1]
		lastSend := time.Unix(currentEvent.Provider.Timestamp, 0)
		nextSend := time.Unix(nextEvent.Provider.Timestamp, 0)
		payload, err := json.MarshalIndent(currentEvent, "", "  ")
		if err != nil {
			fmt.Println("Error encoding")
		}
		fmt.Println("Sending event:", currentEvent, string(payload))
		_, err = http.Post(receiver, "application/json", bytes.NewReader(payload))
		if err != nil {
			fmt.Println("Error posting:", err)
		}
		sleep := nextSend.Sub(lastSend)
		log.Println("Sleeping for: ", sleep)
		time.Sleep(sleep)
	}
}

func main() {
	options := parseArgs()
	if *options.Server == "" {
		fmt.Println("Must supply receiving server.")
		os.Exit(NO_GSI_RECEIVER)
	}
	data, err := ioutil.ReadFile(*options.File)
	if err != nil {
		fmt.Println("Could not read events file: ", *options.File)
		os.Exit(FILE_NOT_READ)
	}
	var gsidata []csgogsi.GSIData
	json.Unmarshal(data, &gsidata)
	if err != nil {
		fmt.Println(err)
		os.Exit(NO_GSI_RECEIVER)
	}
	repeat(gsidata, *options.Server)
}
