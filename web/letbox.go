package web

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"net/http/cookiejar"
)

type Letbox struct {
}

func (l Letbox) Crawl() {
	j, _ := cookiejar.New(nil)
	client := &http.Client{Jar: j}
	req, err := http.NewRequest("GET", "https://my.letbox.com/cart.php?a=add&pid=67", nil)
	checkError(err)
	res, err := client.Do(req)

	checkError(err)
	defer res.Body.Close()
	dom, err := goquery.NewDocumentFromReader(res.Body)
	checkError(err)
	s := dom.Find("h1").Text()

	if s != "Out of Stock" {
		notice("letbox")
	}
	log.Println("letbox finshed!!!")
}
