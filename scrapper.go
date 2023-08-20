package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/silverhand7/go-rss-aggregator/internal/database"
)

func startScrapping(
	db *database.Queries,
	concurrency int, // how many goroutines we want to use to go fetch all of these different feeds (fetch them all at the same time)
	timeBetweenRequest time.Duration,
) {
	log.Printf("Scrapping on %v goroutines every %s duration", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)

	// execute the body of the for loops everytime a new value come across ticker channel.
	// the ticker has a field called C which is a channel where time between request would be sent across the channel
	// for example if we set the timeBetweenRequest to one minute, then every one minute a value would be sent across the channel
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(
			context.Background(),
			int32(concurrency),
		)
		if err != nil {
			log.Println("error fetching feeds:", err)
			// even if there's an error the function keeps running (not stopping completely)
			continue
		}

		// HERE'S MORE EXPLANATION ABOUT THIS CODE: https://youtu.be/un6ZyFkqFKo?t=32313 (make sure to check)

		waitGroup := &sync.WaitGroup{}
		for _, feed := range feeds {
			// we're iterating all of the feeds on the same goroutine
			// we're adding 1 to the wait group for every feed
			waitGroup.Add(1)

			// we'll be spawning all of these separate goroutines, when we get to the end of the loop, we're gonna be waiting on the wait group for distinct call (panggilan yang berbeda) to waitGroup.Done().
			// waitGroup.Done automatically decrement the counter by one
			go scrapeFeed(db, feed, waitGroup)
		}
		waitGroup.Wait()

	}
}

func scrapeFeed(db *database.Queries, feed database.Feed, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error marking feed as fetched:", err)
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println("Error fetching feed:", err)
	}

	for _, item := range rssFeed.Channel.Items {
		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}
		// because the time is using RFC1123Z format
		publishedAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("couldn't parse date %v with err %v", item.PubDate, err)
		}
		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Description: description,
			PublishedAt: sql.NullTime{
				Time:  publishedAt,
				Valid: true,
			},
			Url:    item.Link,
			FeedID: feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Println("failed to create post: ", err)
		}
	}

	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Items))
}
