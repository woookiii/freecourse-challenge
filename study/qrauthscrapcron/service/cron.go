package service

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"qrauthscrapcron/repository"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/robfig/cron"
)

type cronJob struct {
	repository repository.RepositoryImpl
	c          *cron.Cron
}

func NewCronJob(repository repository.RepositoryImpl) *cronJob {
	c := &cronJob{repository: repository, c: cron.New()}

	go c.runJobs()

	return c
}

func (cj *cronJob) runJobs() {
	c := cj.c
	db := cj.repository

	c.AddFunc("*/5 * * * * *", func() {
		cj.scrapping(db)
	})

	c.Start()
	defer c.Stop()

	select {}
}

func (cj *cronJob) scrapping(db repository.RepositoryImpl) error {
	log.Println("five seconds job executed from mysql for scrapping")
	if result, err := db.ViewAll(); err != nil {
		return err
	} else if len(result) == 0 {
		return errors.New("there is no result")
	} else {
		for _, r := range result {
			log.Printf("Try Scrapping URL : %s", r.URL)
			log.Printf("Try Scrapping CardSelect : %s", r.CardSelector)
			log.Printf("Try Scrapping Tag : %s", r.Tag)
			cj.scrappingHTML(r.URL, r.CardSelector, r.InnerSelector, strings.Split(r.Tag, " "))
		}
		return nil
	}
}

func (cj *cronJob) scrappingHTML(url, cardSelector, innerSelector string, tag []string) {
	client := http.Client{Timeout: time.Second * 3}
	if request, err := http.NewRequest("GET", url, nil); err != nil {
		log.Println("Failed to make request", "err", err)
	} else {
		request.Header.Set("User-Agent", "M")

		if response, err := client.Do(request); err != nil {
			log.Println("Failed to call get api", "err", err)
		} else {
			defer response.Body.Close()

			if doc, err := goquery.NewDocumentFromReader(response.Body); err != nil {
				log.Println("Failed to read response", "err", err)
			} else {
				searchCard := doc.Find(cardSelector)

				if searchCard.Length() == 0 {
					log.Println("nothing in CardSelector")
				} else {
					searchCard.Each(func(_ int, card *goquery.Selection) {
						card.Find(innerSelector).Each(func(_ int, child *goquery.Selection) {
							for _, t := range tag {
								d := child.Find(t).Text()
								log.Println(d)
							}
						})
					})
				}
				fmt.Println(doc.Html())
			}
		}
	}
}
