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


}
