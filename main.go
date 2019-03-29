package main

import (
	"fmt"
	"log"

	"github.com/sclevine/agouti"
)

func main() {
	go serve()
	url := "http://localhost:3000/"
	capabilities := agouti.NewCapabilities()
	// capabilities["phantomjs.binary.path"] = "/opt/tweet-style-graphic-generator/node_modules/phantomjs/bin/phantomjs"
	capabilities["phantomjs.binary.path"] = "."
	capabilitiesOption := agouti.Desired(capabilities)
	driver := agouti.PhantomJS(capabilitiesOption)

	err := driver.Start()
	if err != nil {
		log.Fatal("Failed to start driver: %v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage(agouti.Browser("phantomjs"))
	if err != nil {
		log.Fatal("Failed to open page: %v", err)
	}

	err = page.Navigate(url)
	if err != nil {
		log.Fatal("Failed to navigate: %v", err)
	}

	err = page.Screenshot("page.png")
	if err != nil {
		log.Fatal(err)
	}

	html, err := page.HTML()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(html)

}
