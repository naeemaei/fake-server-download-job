package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/naeemaei/fake-server-download-job/jobs"
)

func main() {
	log.Printf("Start of application")
	s := gocron.NewScheduler(time.UTC)
	gocron.SetPanicHandler(func(jobName string, _ any) {
		fmt.Printf("Panic in job: %s", jobName)
		fmt.Println("do something to handle the panic")
	})
	s.SingletonModeAll()
	job, err := s.Every(2).Second().Do(jobs.RunCrawler)
	if err != nil {
		log.Printf("Error creating cron-> id: %v, err: %s", job, err.Error())
	}
	s.StartBlocking()
	log.Printf("End of application")
}
