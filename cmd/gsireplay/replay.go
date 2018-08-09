package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/bergmannf/gsi-simul/csgogsi"
	"io/ioutil"
)

const GSI_FILE = "./gsi.json"

type options struct {
	File   *string
	Config *string
}

func parseArgs() options {
	var opts options
	opts.File = flag.String("file", GSI_FILE, "File to replay events from.")
	opts.Config = flag.String("config", "", "Configuration of the GSI based application.")
	flag.Parse()
	return opts
}

func repeater(eventsFile string, configurationFile string) {
}

func main() {
	options := parseArgs()
	if *options.Config != "" {
		var config csgogsi.GSIConfiguration
		data, _ := ioutil.ReadFile(*options.Config)
		json.Unmarshal(data, &config)
	}
	data, _ := ioutil.ReadFile("./single.json")
	var gsidata csgogsi.GSIData
	json.Unmarshal(data, &gsidata)
	fmt.Printf("%#v", gsidata)
}
