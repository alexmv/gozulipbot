package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	gzb "github.com/ifo/gozulipbot"
)

func main() {
	emailAddress, apiKey, err := gzb.GetConfigFromFlags()
	if err != nil {
		log.Fatalln(err)
	}

	bot := gzb.MakeBot(emailAddress, apiKey, []string{"test-bot"})

	/*
		resp, err := bot.SendStreamMessage("test-bot", "test-go-bot",
			"continued progress")
	*/

	resp, err := bot.GetStreamList()
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var toPrint bytes.Buffer

	err = json.Indent(&toPrint, body, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(toPrint.String())
}
