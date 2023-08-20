package main

import (
	"log"
	"time"

	"github.com/silverhand7/go-rss-aggregator/internal/database"
)

func startScrapping(
	db *database.Queries,
	concurrency int,
	timeBetweenRequest time.Duration,
) {
	log.Printf("Scrapping on %v goroutines every %s duration", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)

	// execute the body of the for loops everytime a new value come across ticker channel.
	// the ticker has a field called C which is a channel where time between request would be sent across the channel
	// for example if we set the timeBetweenRequest to one minute, then every one minute a value would be sent across the channel
	for ; ; <-ticker.C {

	}
}
