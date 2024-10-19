package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/vikram-parashar/rss-rush/database"
)

func startScrapping(numAtOnce int, interval time.Duration, db *database.Queries) {
	ticker := time.NewTicker(interval)

	for {
		channels, err := db.GetChannels(context.Background(), database.GetChannelsParams{
			Limit:  int32(numAtOnce),
			Offset: 0,
		})
		if err != nil {
			log.Printf("couldnt fetch channlels from db: %v", err)
			return
		}

		wg := &sync.WaitGroup{}
		for _, channel := range channels {
			wg.Add(1)
			xlm_url := channel.XmlUrl
			go scrapeUrl(xlm_url, db, channel.ID, wg)
		}
		wg.Wait()
    log.Println("batch done")

		<-ticker.C
	}

}

type Page struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Items       []Article `xml:"item"`
	} `xml:"channel"`
}

type Article struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func scrapeUrl(url string, db *database.Queries, channelId uuid.UUID, wg *sync.WaitGroup) {
	defer wg.Done()

  err:=db.UpdateFetched(context.Background(),channelId)
  if err!=nil{
    log.Printf("fetch time cant be updated: %v", err)
		return
  }

	res, err := http.Get(url)
	if err != nil {
		log.Printf("url cant be fetched %v", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("cant read response: %v", err)
		return
	}

	var page Page
	err = xml.Unmarshal(body, &page)

	articles := page.Channel.Items

	for _, article := range articles {
		pubAt, err := time.Parse(time.RFC1123, article.PubDate)
		if err != nil {
			log.Printf("pub date cant be parsed: %v", err)
			continue
		}

		res, err := db.CreateArticle(context.Background(), database.CreateArticleParams{
			Title: article.Title,
			Description: sql.NullString{
				String: article.Description,
				Valid:  article.Description != "",
			},
			Link:      article.Link,
			PubDate:   pubAt,
			ChannelID: channelId,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate") {
				continue
			}
			log.Printf("couldnt insert article with url %v: %v", article.Link, err)
		}
		log.Printf("Article inserted: %v", res.Title)
	}

}
