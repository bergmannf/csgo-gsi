package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const GSI_FILE = "./gsi.json"

type options struct {
	File *string
	Port *int
}

func parseArgs() options {
	var opts options
	opts.File = flag.String("file", GSI_FILE, "File to store the events to.")
	opts.Port = flag.Int("port", 8080, "Port to listen for events")
	flag.Parse()
	return opts
}

func store(file *os.File) func(http.ResponseWriter, *http.Request) {
	processedEvents := 0
	return func(writer http.ResponseWriter, request *http.Request) {
		body, err := ioutil.ReadAll(request.Body)
		request.Body.Close()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Could not read request body.")
		}
		// Only write on new events and not before the first one.
		// Otherwise the json syntax is invalid: [,{}]
		if processedEvents > 0 {
			file.WriteString(",")
		}
		fmt.Println("Received event: ", processedEvents)
		file.Write(body)
		processedEvents++
	}
}

func logError(err error) {
	if err != nil {
		fmt.Println("Error during writing: ", err)
	}
}

func main() {
	jsonStart := []byte("[")
	jsonEnd := []byte("]")
	options := parseArgs()
	file, err := os.OpenFile(*options.File, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer file.Close()
	logError(err)
	http.HandleFunc("/", store(file))
	_, err = file.Write(jsonStart)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		_, err = file.Write(jsonEnd)
		logError(err)
		os.Exit(1)
	}()

	logError(err)
	fmt.Println("Writing to: ", *options.File)
	fmt.Println("Listening on port: ", *options.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", *options.Port), nil)
}
