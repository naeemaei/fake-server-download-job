package jobs

import (
	"log"
	"math/rand"
	"time"

	"github.com/naeemaei/fake-server-download-job/utility"
	ptime "github.com/yaa110/go-persian-calendar"
)

var urls = []string{
	"http://www.varzesh3.com",
	"http://divar.ir",
	"http://ikco.ir",
	"http://eitaa.com",
	"http://rubika.ir/",
	"http://bale.ai",
	"http://emalls.ir",
	"http://emofid.com",
	"http://yektanet.com",
	"http://zhaket.com",
	"http://isna.ir",
	"http://jobinja.ir",
	"http://time.ir",
	"http://netafraz.com",
	"http://khabaronline.ir",
	"http://blogfa.com",
	"http://faradars.org",
	"http://irna.ir",
	"http://agah.com",
	"http://bama.ir",
	"http://alibaba.ir",
	"http://hamshahrionline.ir",
	"http://niazerooz.com",
}

var httpClient = utility.HttpClient()

// Run from 2 to 10 morning
func canExecute() bool {
	pt := ptime.Now()
	if pt.Hour() >= 0 && pt.Hour() <= 24 {
		return true
	}
	return false
}

func RunCrawler() {
	if canExecute() {
		log.Printf("Start of crawl")
		rand.Seed(time.Now().UnixNano())
		num := rand.Intn(len(urls) - 1)
		utility.SendRequest(httpClient, urls[num], "GET")
		log.Printf("End of crawl")
		return
	}
	log.Printf("Crawler is sleeping")
}
