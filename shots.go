package main

import (
	"log"

	"github.com/sclevine/agouti"
)

func shots() {
	url := "http://localhost:3000/"
	capabilities := agouti.NewCapabilities()
	// capabilities["phantomjs.binary.path"] = "/opt/tweet-style-graphic-generator/node_modules/phantomjs/bin/phantomjs"
	capabilities["phantomjs.binary.path"] = "."
	capabilitiesOption := agouti.Desired(capabilities)
	driver := agouti.PhantomJS(capabilitiesOption)

	err := driver.Start()
	if err != nil {
		log.Fatal(err)
	}
	defer driver.Stop()

	page, err := driver.NewPage(agouti.Browser("phantomjs"))
	if err != nil {
		log.Fatal(err)
	}
	page.Size(1280, 720)

	err = page.Navigate(url)
	if err != nil {
		log.Fatal(err)
	}

	err = page.Screenshot("page.png")
	if err != nil {
		log.Fatal(err)
	}

}
